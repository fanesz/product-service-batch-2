package entity

import "codebase-app/pkg/types"

type ProductsResponse struct {
	Items []ProductItem `json:"items"`
	Meta  types.Meta    `json:"meta"`
}

type ProductItem struct {
	Id          string  `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	Stock       int     `json:"stock" db:"stock"`
	CategoryId  *string `json:"category_id" db:"category_id"`
}

type GetProductResponse struct {
	Name        string              `json:"name" db:"name"`
	Description string              `json:"description" db:"description"`
	Price       float64             `json:"price" db:"price"`
	Stock       int                 `json:"stock" db:"stock"`
	CategoryId  *string             `json:"category_id,omitempty" db:"category_id"`
	Category    GetCategoryResponse `json:"category" db:"category"`
}

type GetCategoryResponse struct {
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

type CreateProductRequest struct {
	ShopId string `json:"shop_id" validate:"uuid" db:"shop_id"`

	Name        string  `json:"name" validate:"required" db:"name"`
	Description string  `json:"description" validate:"required,max=255" db:"description"`
	Price       float64 `json:"price" validate:"required" db:"price"`
	Stock       int     `json:"stock" validate:"required" db:"stock"`
	CategoryId  string  `json:"category_id" db:"category_id"`
}

type CreateProductResponse struct {
	Id string `json:"id" db:"id"`
}

type UpdateProductRequest struct {
	Id          string  `params:"id" validate:"uuid" db:"id"`
	Name        string  `json:"name" validate:"required" db:"name"`
	Description string  `json:"description" validate:"required" db:"description"`
	Price       float64 `json:"price" validate:"required" db:"price"`
	Stock       int     `json:"stock" validate:"required" db:"stock"`
	CategoryId  string  `json:"category_id" db:"category_id"`
}

type UpdateProductResponse struct {
	Id string `json:"id" db:"id"`
}

type DeleteProductRequest struct {
	Id string `validate:"uuid" db:"id"`
}

type GetProductRequest struct {
	Id string `validate:"uuid" db:"id"`
}

type ProductsRequest struct {
	Page     int `query:"page" validate:"required"`
	Paginate int `query:"paginate" validate:"required"`
}

func (r *ProductsRequest) SetDefault() {
	if r.Page < 1 {
		r.Page = 1
	}

	if r.Paginate < 1 {
		r.Paginate = 10
	}
}
