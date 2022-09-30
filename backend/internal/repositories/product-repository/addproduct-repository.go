package product_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
)

// Add adds a list of products in database
func (r *productRepository) Add(ctx context.Context, p []entities.Product) *customerror.CustomError {
	query := "INSERT INTO products(name, producer_id) VALUES(:name, :producer_id) ON CONFLICT (name, producer_id) DO NOTHING"
	_, err := r.sqlx.NamedQueryContext(ctx, query, p)
	if err != nil {
		return customerror.NewError(customerror.EINVALID, "Error", "producer_repository.AddProduct", err)
	}
	return nil
}
