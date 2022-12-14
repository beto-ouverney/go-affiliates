package product_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	"github.com/jmoiron/sqlx"
)

// IProductRepository presents product repository interface
type IProductRepository interface {
	Add(ctx context.Context, p []entities.Product) *customerror.CustomError
	GetAll(ctx context.Context) (*[]entities.Product, *customerror.CustomError)
}

type productRepository struct {
	sqlx *sqlx.DB
}

// New creates a new product repository
func New(db *sqlx.DB) IProductRepository {
	return &productRepository{
		db,
	}
}
