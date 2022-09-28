package tests_test

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	mocks_product_repository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/product-repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_productRepository_AddProduct(t *testing.T) {
	assertions := assert.New(t)

	idMock := int64(1)

	type args struct {
		name        string
		producer_id int64
	}

	tests := []struct {
		describe string
		args     args
		want     *int64
		want1    *customerror.CustomError
		msg      string
		msg1     string
	}{
		{
			describe: "Should be able add a product and return the id",
			args: args{
				name:        "Hubla best company to contents producers",
				producer_id: 1,
			},
			want:  &idMock,
			want1: nil,
			msg:   "The id should be 1",
			msg1:  "The error should be nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			ctx := context.Background()
			m := new(mocks_product_repository.IProductRepository)
			m.On("AddProduct", mock.AnythingOfType("*context.emptyCtx"), tt.args.name, tt.args.producer_id).Return(tt.want, tt.want1)

			got, got1 := m.AddProduct(ctx, tt.args.name, tt.args.producer_id)
			assertions.Equal(tt.want, got, tt.msg)
			assertions.Equal(tt.want1, got1, tt.msg1)

		})
	}
}
