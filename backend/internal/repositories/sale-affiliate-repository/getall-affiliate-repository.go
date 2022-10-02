package sale_affiliate_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
)

// GetAll returns all sales affiliates with name of product, producer and affiliate
func (r *saleAffiliateRepository) GetAll(ctx context.Context) (*[]entities.SaleResponse, *customerror.CustomError) {

	query := `SELECT sales_affiliates.value, sales_affiliates.commission, sales_affiliates.date, producers.name as "producer", affiliates.name as "affiliate", products.name as "product" FROM sales_affiliates INNER JOIN producers ON sales_affiliates.producer_id = producers.id INNER JOIN affiliates ON sales_affiliates.affiliate_id = affiliates.id INNER JOIN products ON sales_affiliates.product_id = products.id`

	sales := []entities.SaleResponse{}
	err := r.sqlx.SelectContext(ctx, &sales, query)
	if err != nil {
		return nil, customerror.NewError(customerror.EINVALID, "Error", "sale_repository.GetAll", err)
	}

	return &sales, nil
}
