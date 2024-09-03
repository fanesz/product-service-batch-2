package model

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID  `json:"id"`
	UserID    uuid.UUID  `json:"user_id"`
	ProductID uuid.UUID  `json:"product_id"`
	Qty       int        `json:"qty"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type GetCartRequest struct {
	UserID    uuid.UUID   `json:"user_id"`
	ProductID []uuid.UUID `json:"product_id"`
}

type UpdateCartRequest struct {
	UserID uuid.UUID            `json:"user_id"`
	Items  []ProductItemRequest `json:"items"`
}

type ProductItemRequest struct {
	ProductID uuid.UUID `json:"product_id"`
	Qty       int       `json:"qty"`
}

type DeleteCartRequest struct {
	UserID    uuid.UUID `json:"user_id"`
	ProductID uuid.UUID `json:"product_id"`
}
