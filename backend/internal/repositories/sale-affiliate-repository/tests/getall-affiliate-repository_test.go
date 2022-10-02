package tests

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	mocksaffiliaterepository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/sale-affiliate-repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_saleAffiliateRepository_GetAll(t *testing.T) {
	assertions := assert.New(t)

	mockSales := []entities.SaleResponse{
		{
			Producer:   "Producer 1",
			Affiliate:  "Affiliate 1",
			Product:    "Product 1",
			Value:      1000,
			Commission: 100,
			Date:       "22020-02-03T20:51:59-03:00",
		},
		{
			Producer:   "Producer 2",
			Affiliate:  "Affiliate 2",
			Product:    "Product 2",
			Value:      2000,
			Commission: 500,
			Date:       "22022-02-03T20:51:59-03:00",
		},
	}

	tests := []struct {
		describe string
		args     []entities.SaleResponse
		want     *[]entities.SaleResponse
		want1    *customerror.CustomError
		msg      string
		msg1     string
	}{
		{
			describe: "Should be able get all sales with products, sales and affiliates",
			args:     mockSales,
			want:     &mockSales,
			want1:    nil,
			msg:      "Must be equal",
			msg1:     "The error should be nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			ctx := context.Background()
			m := new(mocksaffiliaterepository.ISaleAffiliateRepository)
			m.On("GetAll", mock.AnythingOfType("*context.emptyCtx")).Return(tt.want, tt.want1)

			got, got1 := m.GetAll(ctx)
			assertions.Equal(tt.want1, got1, tt.msg1)
			assertions.EqualValues(tt.want, got, tt.msg)

		})
	}
}
