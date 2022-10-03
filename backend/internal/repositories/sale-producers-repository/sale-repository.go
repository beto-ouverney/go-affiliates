package sale_producers_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	"github.com/jmoiron/sqlx"
)

// ISaleRepository presents the interface for the sales repository
type ISaleRepository interface {
	Add(ctx context.Context, s []entities.Sale) *customerror.CustomError
	GetAll(ctx context.Context) (*[]entities.SaleResponse, *customerror.CustomError)
}

type saleRepository struct {
	sqlx *sqlx.DB
}

// New creates a new sales repository
func New(db *sqlx.DB) ISaleRepository {
	return &saleRepository{
		db,
	}
}
