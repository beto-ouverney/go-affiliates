package producer_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
)

// GetByName gets a content producer by name in database
func (r *producerRepository) GetByName(ctx context.Context, name string) (*entities.Producer, *customerror.CustomError) {
	var producer entities.Producer
	err := r.sqlx.GetContext(ctx, &producer, "SELECT id, name FROM producers WHERE name = $1", name)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, customerror.NewError(customerror.ENOTFOUND, "Error", "producer_repository.GetByName", err)
		}
		return nil, customerror.NewError(customerror.EINVALID, "Error", "producer_repository.GetByName", err)
	}
	return &producer, nil
}
