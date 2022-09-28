package affiliate_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
)

// GetByName gets an affiliate by name in database
func (r *affiliateRepository) GetByName(ctx context.Context, name string) (*entities.Affiliate, *customerror.CustomError) {
	var affiliate entities.Affiliate
	err := r.sqlx.GetContext(ctx, &affiliate, "SELECT id, name FROM affiliates WHERE name = $1", name)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, customerror.NewError(customerror.ENOTFOUND, "Not found", "affiliate_repository.GetByName", err)
		}
		return nil, customerror.NewError(customerror.EINVALID, "Error", "affiliate_repository.GetByName", err)
	}
	return &affiliate, nil
}
