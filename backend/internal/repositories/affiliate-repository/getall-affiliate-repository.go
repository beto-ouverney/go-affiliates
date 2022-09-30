package affiliate_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
)

// GetAll gets all affiliates from database
func (r *affiliateRepository) GetAll(ctx context.Context) (*[]entities.Affiliate, *customerror.CustomError) {
	var af []entities.Affiliate
	err := r.sqlx.SelectContext(ctx, &af, "SELECT id, name, producer_id FROM affiliates")
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, customerror.NewError(customerror.ENOTFOUND, "Not found", "affiliate_repository.GetAll", err)
		}
		return nil, customerror.NewError(customerror.EINVALID, "Error", "affiliate_repository.GetAll", err)
	}

	return &af, nil
}
