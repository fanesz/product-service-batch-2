package ports

import (
	"codebase-app/internal/module/product/entity"
	"context"
)

type ProductRepository interface {
	GetProduct(ctx context.Context, req *entity.GetProductRequest) (*entity.GetProductResponse, error)
	GetProducts(ctx context.Context, req *entity.ProductsRequest) (*entity.ProductsResponse, error)
	CreateProduct(ctx context.Context, req *entity.CreateProductRequest) (*entity.CreateProductResponse, error)
	UpdateProduct(ctx context.Context, req *entity.UpdateProductRequest) (*entity.UpdateProductResponse, error)
	DeleteProduct(ctx context.Context, req *entity.DeleteProductRequest) error
}

type ProductService interface {
	GetProduct(ctx context.Context, req *entity.GetProductRequest) (*entity.GetProductResponse, error)
	GetProducts(ctx context.Context, req *entity.ProductsRequest) (*entity.ProductsResponse, error)
	CreateProduct(ctx context.Context, req *entity.CreateProductRequest) (*entity.CreateProductResponse, error)
	UpdateProduct(ctx context.Context, req *entity.UpdateProductRequest) (*entity.UpdateProductResponse, error)
	DeleteProduct(ctx context.Context, req *entity.DeleteProductRequest) error
}
