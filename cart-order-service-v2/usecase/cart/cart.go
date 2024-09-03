package cart

import (
	model "cart-order-service/repository/models"
	"fmt"
)

// cartStore is an interface that defines the methods required for managing a shopping cart.
type cartStore interface {
	// GetCartByUserID retrieves the cart for a given user.
	GetCartByUserID(bReq model.GetCartRequest) (*[]model.Cart, error)
	// UpdateCartByUserID updates the cart for a given user.
	UpdateCartByUserID(bReq model.UpdateCartRequest) (*[]model.Cart, error)
}

// cart is a struct that holds the store for managing a shopping cart.
type cart struct {
	store cartStore
}

// NewCart is a constructor function that returns a new cart instance.
func NewCart(store cartStore) *cart {
	return &cart{store}
}

// GetCartByUserID is a method that retrieves the cart for a given user and returns a response with the total items.
func (c *cart) GetCartByUserID(bReq model.GetCartRequest) (*[]model.Cart, error) {
	result, err := c.store.GetCartByUserID(bReq)
	if err != nil {
		return nil, err
	}

	if len(*result) == 0 {
		return nil, nil
	}

	return result, nil
}

func (c *cart) UpdateCartByUserID(bReq model.UpdateCartRequest) (*[]model.Cart, error) {
	result, err := c.store.UpdateCartByUserID(bReq)
	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	fmt.Println(result)
	return result, nil
}
