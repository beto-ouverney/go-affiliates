package sale_affiliate_repository

import (
	"context"
	"fmt"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
)

// Add a affiliate sale or more in database
func (r *saleAffiliateRepository) Add(ctx context.Context, sale []entities.Sale) *customerror.CustomError {
	fmt.Println(sale)
	query := "INSERT INTO sales_affiliates (product_id, producer_id,affiliate_id, value, commission, date ) VALUES (:product_id, :producer_id,:affiliate_id, :value, :commission, :date)  ON CONFLICT (product_id,producer_id,affiliate_id, date) DO NOTHING"
	rows, err := r.sqlx.NamedQueryContext(ctx, query, sale)
	if err != nil {
		fmt.Println(err)
		return customerror.NewError(customerror.EINVALID, "Error", "sale_repository.Add", err)
	}

	fmt.Println(*rows)
	return nil
}
