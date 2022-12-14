package producer_repository

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	"github.com/jmoiron/sqlx"
)

// IProducerRepository presents producer repository interface
type IProducerRepository interface {
	Add(ctx context.Context, p []entities.Producer) *customerror.CustomError
	GetAll(ctx context.Context) (*[]entities.Producer, *customerror.CustomError)
}

type producerRepository struct {
	sqlx *sqlx.DB
}

// New creates a new producer repository
func New(db *sqlx.DB) IProducerRepository {
	return &producerRepository{
		db,
	}
}
