package tests

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/controllers/sales-controller/mocks"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_saleController_GetAll(t *testing.T) {
	assertions := assert.New(t)

	mockSales := []entities.SaleResponse{
		{
			Producer:   "Producer 1",
			Product:    "Product 1",
			Value:      1000,
			Commission: 0,
			Date:       "22020-02-03T20:51:59-03:00",
		},
		{
			Producer:   "Producer 2",
			Product:    "Product 2",
			Value:      2000,
			Commission: 0,
			Date:       "22022-02-03T20:51:59-03:00",
		},
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
		want     *[]entities.SaleResponse
		want1    *customerror.CustomError
	}{
		{
			describe: "Should be able get all sales",
			want:     &mockSales,
			want1:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			ctx := context.Background()

			m := new(mocks.ISaleController)
			m.On("GetAll", mock.AnythingOfType("*context.emptyCtx")).Return(tt.want, tt.want1)

			got, got1 := m.GetAll(ctx)
			assertions.EqualValues(tt.want, got)
			assertions.Equal(tt.want1, got1)
		})
	}
}
