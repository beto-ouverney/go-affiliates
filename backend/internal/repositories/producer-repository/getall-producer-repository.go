package producer_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
)

// GetAll gets all content producers from database
func (r *producerRepository) GetAll(ctx context.Context) (*[]entities.Producer, *customerror.CustomError) {
	var p []entities.Producer
	err := r.sqlx.SelectContext(ctx, &p, "SELECT id, name FROM producers")
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, customerror.NewError(customerror.ENOTFOUND, "Error", "producer_repository.GetAll", err)
		}
		return nil, customerror.NewError(customerror.EINVALID, "Error", "producer_repository.GetAll", err)
	}

	return &p, nil
}
