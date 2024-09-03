package cart

import (
	"cart-order-service/helper"
	model "cart-order-service/repository/models"
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type store struct {
	db *sql.DB
}

// NewStore is a constructor function that returns a new store instance.
func NewStore(db *sql.DB) *store {
	return &store{db}
}

// GetCartByUserID is a method that retrieves the cart for a given user.
// It returns a slice of cart and an error if any occurs during the retrieval process.
func (s *store) GetCartByUserID(bReq model.GetCartRequest) (*[]model.Cart, error) {
	querySelect := `
		SELECT
			*
		FROM cart_items
	`

	var queryConditions []string

	if bReq.UserID != uuid.Nil {
		queryConditions = append(queryConditions, fmt.Sprintf("user_id = '%s'", bReq.UserID))
	}

	if len(bReq.ProductID) > 0 {
		var productIDs []string
		for _, pid := range bReq.ProductID {
			productIDs = append(productIDs, fmt.Sprintf("'%s'", pid))
		}
		queryConditions = append(queryConditions, fmt.Sprintf("product_id IN (%s)", strings.Join(productIDs, ",")))
	}

	if len(queryConditions) > 0 {
		querySelect += " WHERE " + strings.Join(queryConditions, " AND ")
	} else {
		querySelect += " WHERE deleted_at IS NULL"
	}

	querySelect += " AND deleted_at IS NULL"

	rows, err := s.db.Query(querySelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var carts []model.Cart
	for rows.Next() {
		var cart model.Cart
		if err := rows.Scan(
			&cart.ID,
			&cart.UserID,
			&cart.ProductID,
			&cart.Qty,
			&cart.CreatedAt,
			&cart.UpdatedAt,
			&cart.DeletedAt,
		); err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &carts, nil
}

func (s *store) UpdateCartByUserID(bReq model.UpdateCartRequest) (*[]model.Cart, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	itemsMap := make(map[uuid.UUID]int)
	for _, item := range bReq.Items {
		itemsMap[item.ProductID] = item.Qty
	}

	for productID, qty := range itemsMap {
		_, err := tx.Exec(`
			INSERT INTO cart_items (user_id, product_id, qty)
			VALUES ($1, $2, $3)
			ON CONFLICT (user_id, product_id)
			DO UPDATE SET qty = EXCLUDED.qty
		`, bReq.UserID, productID, qty)
		if err != nil {
			return nil, err
		}
	}

	productIDs := make([]interface{}, len(itemsMap))
	i := 0
	for productID := range itemsMap {
		productIDs[i] = productID
		i++
	}

	if len(productIDs) > 0 {
		query := `
			DELETE FROM cart_items
			WHERE user_id = $1 AND product_id NOT IN (` + helper.Placeholders(len(productIDs), 2) + `)
		`
		args := append([]interface{}{bReq.UserID}, productIDs...)
		_, err = tx.Exec(query, args...)
		if err != nil {
			return nil, err
		}
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	return s.GetCartByUserID(model.GetCartRequest{UserID: bReq.UserID})
}
