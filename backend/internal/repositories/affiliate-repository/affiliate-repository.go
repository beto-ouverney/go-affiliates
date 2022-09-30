package affiliate_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/db"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	"github.com/jmoiron/sqlx"
)

// IAffiliateRepository presents the interface for the affiliate repository
type IAffiliateRepository interface {
	Add(ctx context.Context, af []entities.Affiliate) *customerror.CustomError
	GetAll(ctx context.Context) (*[]entities.Affiliate, *customerror.CustomError)
}

type affiliateRepository struct {
	sqlx *sqlx.DB
}

// New creates a new affiliate repository
func New() IAffiliateRepository {
	return &affiliateRepository{
		db.ConnectDB(),
	}
}
