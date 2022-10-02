package sale_affiliate_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/db"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	"github.com/jmoiron/sqlx"
)

// ISaleAffiliateRepository presents the interface for the sales repository
type ISaleAffiliateRepository interface {
	Add(ctx context.Context, s []entities.Sale) *customerror.CustomError
	GetAll(ctx context.Context) (*[]entities.SaleResponse, *customerror.CustomError)
}

type saleAffiliateRepository struct {
	sqlx *sqlx.DB
}

// New creates a new sales repository
func New() ISaleAffiliateRepository {
	return &saleAffiliateRepository{
		db.ConnectDB(),
	}
}
