package affiliate_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
)

// Add adds a new affiliate in database
func (r *affiliateRepository) Add(ctx context.Context, name string, product_id int64) (*int64, *customerror.CustomError) {
	var id int64
	err := r.sqlx.GetContext(ctx, &id, "INSERT INTO affiliates(name) VALUES($1,$2) RETURNING id", name, product_id)
	if err != nil {
		return nil, customerror.NewError(customerror.EINVALID, "Error", "affiliate_repository.Add", err)
	}
	return &id, nil
}
