package repository

import (
	"codebase-app/internal/module/product/entity"
	"codebase-app/internal/module/product/ports"
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

var _ ports.ProductRepository = &productRepository{}

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *productRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) GetProducts(ctx context.Context, req *entity.ProductsRequest) (*entity.ProductsResponse, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.ProductItem
	}

	var (
		resp = new(entity.ProductsResponse)
		data = make([]dao, 0, req.Paginate)
	)
	resp.Items = make([]entity.ProductItem, 0, req.Paginate)

	query := `
		SELECT
			COUNT(id) OVER() as total_data,
			id, name, description, price, stock, category_id
		FROM products
		WHERE
			deleted_at IS NULL
		LIMIT ? OFFSET ?
	`
	err := r.db.SelectContext(ctx, &data, r.db.Rebind(query),
		req.Paginate,
		req.Paginate*(req.Page-1),
	)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::GetProducts - Failed to get products")
		return nil, err
	}

	if len(data) > 0 {
		resp.Meta.TotalData = data[0].TotalData
	}

	for _, d := range data {
		resp.Items = append(resp.Items, d.ProductItem)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}

func (r *productRepository) GetProduct(ctx context.Context, req *entity.GetProductRequest) (*entity.GetProductResponse, error) {
	var (
		resp = new(entity.GetProductResponse)
	)
	query := `
		SELECT
			p.name,
			p.description,
			p.price,
			p.stock,
			c.name AS "category.name",
			c.description AS "category.description"
		FROM 
			products p
		JOIN 
			categories c ON p.category_id = c.id
		WHERE
			p.deleted_at IS NULL
			AND c.deleted_at IS NULL
			AND p.id = ?
	`
	err := r.db.GetContext(ctx, resp, r.db.Rebind(query), req.Id)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::GetProduct - Failed to get product and category")
		return nil, err
	}

	resp.CategoryId = nil
	return resp, nil
}

func (r *productRepository) CreateProduct(ctx context.Context, req *entity.CreateProductRequest) (*entity.CreateProductResponse, error) {
	var resp = new(entity.CreateProductResponse)
	query := `
		INSERT INTO products (shop_id, name, description, price, stock, category_id)
		VALUES (?, ?, ?, ?, ?, ?)
		RETURNING id
	`

	err := r.db.QueryRowContext(ctx, r.db.Rebind(query),
		req.ShopId,
		req.Name,
		req.Description,
		req.Price,
		req.Stock,
		req.CategoryId,
	).Scan(&resp.Id)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::CreateProduct - Failed to create product")
		return nil, err
	}

	return resp, nil
}

func (r *productRepository) UpdateProduct(ctx context.Context, req *entity.UpdateProductRequest) (*entity.UpdateProductResponse, error) {
	var resp = new(entity.UpdateProductResponse)
	query := `
		UPDATE products
		SET name = ?, description = ?, price = ?, stock = ?, category_id = ?
		WHERE id = ?
		RETURNING id
	`

	err := r.db.QueryRowContext(ctx, r.db.Rebind(query),
		req.Name,
		req.Description,
		req.Price,
		req.Stock,
		req.CategoryId,
		req.Id,
	).Scan(&resp.Id)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::UpdateProduct - Failed to update product")
		return nil, err
	}

	return resp, nil
}

func (r *productRepository) DeleteProduct(ctx context.Context, req *entity.DeleteProductRequest) error {
	query := `
		UPDATE products
		SET deleted_at = NOW()
		WHERE id = ?
	`

	_, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::DeleteProduct - Failed to delete product")
		return err
	}

	return nil
}
