package affiliate_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	"github.com/jmoiron/sqlx"
)

// IAffiliateRepository presents the affiliate repository interface
type IAffiliateRepository interface {
	Add(ctx context.Context, af []entities.Affiliate) *customerror.CustomError
	GetAll(ctx context.Context) (*[]entities.Affiliate, *customerror.CustomError)
}

type affiliateRepository struct {
	sqlx *sqlx.DB
}

// New creates a new affiliate repository
func New(db *sqlx.DB) IAffiliateRepository {
	return &affiliateRepository{
		db,
	}
}
