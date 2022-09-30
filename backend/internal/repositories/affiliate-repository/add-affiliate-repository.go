package affiliate_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
)

// Add adds a affiliates list in database
func (r *affiliateRepository) Add(ctx context.Context, af []entities.Affiliate) *customerror.CustomError {
	query := "INSERT INTO affiliates(name,producer_id) VALUES(:name,:producer_id) ON CONFLICT (name, producer_id) DO NOTHING"
	_, err := r.sqlx.NamedQueryContext(ctx, query, af)
	if err != nil {
		return customerror.NewError(customerror.EINVALID, "Error", "affiliate_repository.Add", err)
	}
	return nil

}
