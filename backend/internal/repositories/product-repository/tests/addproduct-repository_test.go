package tests_test

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	mocks_product_repository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/product-repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_productRepository_AddProduct(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		data []entities.Product
	}

	tests := []struct {
		describe string
		args     args
		want     *customerror.CustomError
		msg      string
	}{
		{
			describe: "Should be able add a products list and return a nil error",
			args: args{
				data: []entities.Product{
					{
						Name:       "Hubla",
						ProducerId: 1,
					},
					{
						Name:       "Alberto Paz",
						ProducerId: 2,
					},
				},
			},
			want: nil,
			msg:  "The error must be nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			ctx := context.Background()
			m := new(mocks_product_repository.IProductRepository)
			m.On("Add", mock.AnythingOfType("*context.emptyCtx"), tt.args.data).Return(tt.want)

			got := m.Add(ctx, tt.args.data)
			assertions.Equal(tt.want, got, tt.msg)

		})
	}
}
