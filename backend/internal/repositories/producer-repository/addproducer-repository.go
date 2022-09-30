package producer_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
)

// Add add content oroducers list in database
func (r *producerRepository) Add(ctx context.Context, p []entities.Producer) *customerror.CustomError {

	query := "INSERT INTO producers(name) VALUES(:name) ON CONFLICT (name) DO NOTHING"

	_, err := r.sqlx.NamedQueryContext(ctx, query, p)
	if err != nil {
		return customerror.NewError(customerror.EINVALID, "Error", "producer_repository.AddProducer", err)
	}

	return nil
}
