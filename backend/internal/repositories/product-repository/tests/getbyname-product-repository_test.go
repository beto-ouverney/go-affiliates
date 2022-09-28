package tests

import (
	"context"
	"errors"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	mocks_product_repository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/product-repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_productRepository_GetProductByName(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		name string
	}

	mockProduct := entities.Product{
		ID:   1,
		Name: "Product 1",
	}
	tests := []struct {
		describe string
		args     args
		want     *entities.Product
		want1    *customerror.CustomError
		msg      string
		msg1     string
	}{
		{
			describe: "Should be able to get a product by name",
			args: args{
				name: "Product 1",
			},
			want:  &mockProduct,
			want1: nil,
			msg:   "The product should be Product 1",
			msg1:  "The error should be nil",
		},
		{
			describe: "Should be able return a nil product and a custom error if product don`t exist in database",
			args: args{
				name: "Product 2",
			},
			want: nil,
			want1: customerror.NewError(customerror.ENOTFOUND, "Not found", "_repository.GetProductByName",
				errors.New("sql: no rows in result set")),
			msg:  "The product should be nil",
			msg1: "The error should be a custom error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			ctx := context.Background()
			m := new(mocks_product_repository.IProductRepository)
			m.On("GetProductByName", mock.AnythingOfType("*context.emptyCtx"), tt.args.name).Return(tt.want, tt.want1)

			got, got1 := m.GetProductByName(ctx, tt.args.name)
			assertions.Equal(tt.want, got, tt.msg)
			assertions.Equal(tt.want1, got1, tt.msg1)
		})
	}
}
