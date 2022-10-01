package sale_affiliate_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
)

// Add a affiliate sale or more in database
func (r *saleAffiliateRepository) Add(ctx context.Context, sale []entities.Sale) *customerror.CustomError {

	query := "INSERT INTO sales_affiliates (product_id, producer_id,affiliate_id, value, commission, date ) VALUES (:product_id, :producer_id,:affiliate_id, :value, :commission, :date)  ON CONFLICT (product_id,producer_id,affiliate_id, date) DO NOTHING"
	_, err := r.sqlx.NamedQueryContext(ctx, query, sale)
	if err != nil {
		return customerror.NewError(customerror.EINVALID, "Error", "sale_repository.Add", err)
	}

	return nil
}
