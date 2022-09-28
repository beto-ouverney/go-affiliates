package product_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/db"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	"github.com/jmoiron/sqlx"
)

// IProductRepository presents the interface for the product repository
type IProductRepository interface {
	AddProduct(ctx context.Context, name string, producerId int64) (*int64, *customerror.CustomError)
	GetProductByName(ctx context.Context, name string) (*entities.Product, *customerror.CustomError)
}

type productRepository struct {
	sqlx *sqlx.DB
}

// New creates a new product repository
func New() IProductRepository {
	return &productRepository{
		db.ConnectDB(),
	}
}
