package tests

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	mocks_producer_repository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/producer-repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_producerRepository_Add(t *testing.T) {
	assertions := assert.New(t)

	idMock := int64(1)

	tests := []struct {
		describe string
		args     string
		want     *int64
		want1    *customerror.CustomError
		msg      string
		msg1     string
	}{
		{
			describe: "Should be able add a producer and return the id",
			args:     "Hubla best company to contents producers",
			want:     &idMock,
			want1:    nil,
			msg:      "The id should be 1",
			msg1:     "The error should be nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			ctx := context.Background()
			m := new(mocks_producer_repository.IProducerRepository)
			m.On("Add", mock.AnythingOfType("*context.emptyCtx"), tt.args).Return(tt.want, tt.want1)

			got, got1 := m.Add(ctx, tt.args)
			assertions.Equal(tt.want, got, tt.msg)
			assertions.Equal(tt.want1, got1, tt.msg1)
		})
	}
}
