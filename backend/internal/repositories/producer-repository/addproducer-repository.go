package producer_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
)

// Add adds a new producer in database and return the id
func (r *producerRepository) Add(ctx context.Context, name string) (*int64, *customerror.CustomError) {
	var id int64
	err := r.sqlx.GetContext(ctx, &id, "INSERT INTO producers(name) VALUES($1) RETURNING id", name)
	if err != nil {
		return nil, customerror.NewError(customerror.EINVALID, "Error", "producer_repository.AddProducer", err)
	}
	return &id, nil
}
