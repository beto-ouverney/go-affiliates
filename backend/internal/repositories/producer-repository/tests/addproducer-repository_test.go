package tests

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	mocks_producer_repository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/producer-repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_producerRepository_Add(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		data []entities.Producer
	}

	tests := []struct {
		describe string
		args     args
		want     *customerror.CustomError
		msg      string
	}{
		{
			describe: "Should be able add a producer and return the id",
			args: args{
				data: []entities.Producer{
					{
						Name: "Hubla",
					},
					{
						Name: "Alberto Paz",
					},
				},
			},
			want: nil,
			msg:  "The error should be nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			ctx := context.Background()
			m := new(mocks_producer_repository.IProducerRepository)
			m.On("Add", mock.AnythingOfType("*context.emptyCtx"), tt.args.data).Return(tt.want)

			got := m.Add(ctx, tt.args.data)
			assertions.Equal(tt.want, got, tt.msg)
		})
	}
}
