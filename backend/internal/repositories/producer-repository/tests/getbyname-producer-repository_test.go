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

	type args struct {
		name string
	}

	mockProduct := entities.Producer{
		ID:   1,
		Name: "Alberto Paz",
	}
	tests := []struct {
		describe string
		args     args
		want     *entities.Producer
		want1    *customerror.CustomError
		msg      string
		msg1     string
	}{
		{
			describe: "Should be able to get a content producer by name",
			args: args{
				name: "Alberto Paz",
			},
			want:  &mockProduct,
			want1: nil,
			msg:   "The product should be Alberto Paz",
			msg1:  "The error should be nil",
		},
		{
			describe: "Should be able return a nil producer entity return and a custom error if product don`t exist in database",
			args: args{
				name: "John Doe",
			},
			want: nil,
			want1: customerror.NewError(customerror.ENOTFOUND, "Error", "producer_repository.GetByName",
				errors.New("sql: no rows in result set")),
			msg:  "The product should be nil",
			msg1: "The error should be a custom error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			ctx := context.Background()
			m := new(mocks_producer_repository.IProducerRepository)
			m.On("GetByName", mock.AnythingOfType("*context.emptyCtx"), tt.args.name).Return(tt.want, tt.want1)

			got, got1 := m.GetByName(ctx, tt.args.name)
			assertions.Equal(tt.want, got, tt.msg)
			assertions.Equal(tt.want1, got1, tt.msg1)
		})
	}
}
