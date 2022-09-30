package product_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
)

// GetAll get all products from database
func (r *productRepository) GetAll(ctx context.Context) (*[]entities.Product, *customerror.CustomError) {
	var p []entities.Product
	err := r.sqlx.SelectContext(ctx, &p, "SELECT id, name, producer_id FROM products")
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, customerror.NewError(customerror.ENOTFOUND, "Not found", "_repository.GetAll", err)
		}
		return nil, customerror.NewError(customerror.EINVALID, "Error", "product_repository.GetAll", err)
	}

	return &p, nil
}
