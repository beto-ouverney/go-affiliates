package sale_producers_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
)

// Add producers sales  in database
func (r *saleRepository) Add(ctx context.Context, sales []entities.Sale) *customerror.CustomError {

	query := "INSERT INTO sales_producers (product_id, producer_id, value, commission, date ) VALUES (:product_id, :producer_id, :value, :commission, :date) ON CONFLICT (product_id, producer_id, date) DO NOTHING"
	_, err := r.sqlx.NamedQueryContext(ctx, query, sales)
	if err != nil {
		return customerror.NewError(customerror.EINVALID, "Error", "sale_repository.Add", err)
	}

	return nil
}
