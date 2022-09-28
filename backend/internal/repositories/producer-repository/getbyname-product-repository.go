package producer_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
)

// GetProductByName gets a product by name in database
func (r *productRepository) GetProductByName(ctx context.Context, name string) (*entities.Product, *customerror.CustomError) {
	var product entities.Product
	err := r.sqlx.GetContext(ctx, &product, "SELECT id, name FROM products WHERE name = $1", name)
	if err != nil {
		return nil, customerror.NewError(customerror.EINVALID, "Error", "producer_repository.GetProductByName", err)
	}
	return &product, nil
}
