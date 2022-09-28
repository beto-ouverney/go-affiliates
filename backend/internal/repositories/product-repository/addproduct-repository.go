package product_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
)

// AddProduct adds a new product in database and return the id
func (r *productRepository) AddProduct(ctx context.Context, name string) (*int64, *customerror.CustomError) {
	var id int64
	err := r.sqlx.GetContext(ctx, &id, "INSERT INTO products(name) VALUES($1) RETURNING id", name)
	if err != nil {
		return nil, customerror.NewError(customerror.EINVALID, "Error", "producer_repository.GetAll", err)
	}
	return &id, nil
}
