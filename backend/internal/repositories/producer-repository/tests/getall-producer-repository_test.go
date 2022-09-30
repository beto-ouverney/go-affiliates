package tests

import (
	"context"
	"errors"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	mocks_producer_repository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/producer-repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_producerRepository_GetByName(t *testing.T) {
	assertions := assert.New(t)

	mockProducers := []entities.Producer{
		{
			ID:   1,
			Name: "Alberto Paz",
		},
		{
			ID:   2,
			Name: "Hubla",
		},
	}
	tests := []struct {
		describe string
		want     *[]entities.Producer
		want1    *customerror.CustomError
		msg      string
		msg1     string
	}{
		{
			describe: "Should be able to get all content producers",
			want:     &mockProducers,
			want1:    nil,
			msg:      "The list must be equal",
			msg1:     "The error should be nil",
		},
		{
			describe: "Should be able return a nil list producers return and a custom error if product don`t exist in database",
			want:     nil,
			want1: customerror.NewError(customerror.ENOTFOUND, "Error", "producer_repository.GetByName",
				errors.New("sql: no rows in result set")),
			msg:  "The list must be nil",
			msg1: "The error should be a custom error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			ctx := context.Background()
			m := new(mocks_producer_repository.IProducerRepository)
			m.On("GetAll", mock.AnythingOfType("*context.emptyCtx")).Return(tt.want, tt.want1)

			got, got1 := m.GetAll(ctx)
			assertions.Equal(tt.want, got, tt.msg)
			assertions.Equal(tt.want1, got1, tt.msg1)
		})
	}
}
