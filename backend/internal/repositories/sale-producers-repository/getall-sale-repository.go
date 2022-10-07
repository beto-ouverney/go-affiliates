package sale_producers_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
)

// GetAll returns all sales producers with name of product, producer and affiliate
func (r *saleRepository) GetAll(ctx context.Context) (*[]entities.SaleResponse, *customerror.CustomError) {

	q := `SELECT sales_producers.value, sales_producers.date, producers.name as "producer", products.name as "product" FROM sales_producers INNER JOIN producers ON sales_producers.producer_id = producers.id INNER JOIN products ON sales_producers.product_id = products.id`

	sales := []entities.SaleResponse{}

	err := r.sqlx.SelectContext(ctx, &sales, q)
	if err != nil {
		return nil, customerror.NewError(customerror.EINVALID, "Error", "sale_repository.GetAll", err)
	}

	return &sales, nil
}
