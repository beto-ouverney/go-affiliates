package sale_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
)

// Add a sale or more in database
func (r *saleRepository) Add(ctx context.Context, sale []entities.Sale) *customerror.CustomError {
	query := "INSERT INTO sales (product_id, producer_id, affiliate_id, value, commission, date ) VALUES (:product_id, :producer_id, :affiliate_id, :value, :commission, :date)"
	_, err := r.sqlx.NamedQueryContext(ctx, query, sale)
	if err != nil {
		return customerror.NewError(customerror.EINVALID, "Error", "sale_repository.Add", err)
	}

	return nil
}
