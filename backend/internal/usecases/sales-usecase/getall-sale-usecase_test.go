package sales_usecase

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	mocks_affiliate_repository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/affiliate-repository/mocks"
	mocks_producer_repository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/producer-repository/mocks"
	mocks_product_repository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/product-repository/mocks"
	mocksA "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/sale-affiliate-repository/mocks"
	mocksP "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/sale-producers-repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_salesUseCase_GetAll(t *testing.T) {
	assertions := assert.New(t)

	mockSalesAff := []entities.SaleResponse{
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

	mockSalesPro := []entities.SaleResponse{
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
	}
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
		arg1     *[]entities.SaleResponse
		arg2     *[]entities.SaleResponse
		want     *[]entities.SaleResponse
		want1    *customerror.CustomError
	}{
		{
			describe: "Should be able get all sales",
			arg1:     &mockSalesAff,
			arg2:     &mockSalesPro,
			want:     &mockSales,
			want1:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			ctx := context.Background()
			mockAffRepo := new(mocks_affiliate_repository.IAffiliateRepository)
			mockProducerRepo := new(mocks_producer_repository.IProducerRepository)
			mockProductRepo := new(mocks_product_repository.IProductRepository)
			mockSaleRepo := new(mocksP.ISaleRepository)
			mocSaleAffRepo := new(mocksA.ISaleAffiliateRepository)

			mockSaleRepo.On("GetAll", mock.AnythingOfType("*context.emptyCtx")).Return(tt.arg2, nil)
			mocSaleAffRepo.On("GetAll", mock.AnythingOfType("*context.emptyCtx")).Return(tt.arg1, nil)

			s := salesUseCase{saleRepository: mockSaleRepo, affiliateRepository: mockAffRepo, producerRepository: mockProducerRepo, productRepository: mockProductRepo, saleAffiliateRepository: mocSaleAffRepo}
			got, got1 := s.GetAll(ctx)
			assertions.EqualValues(tt.want, got)
			assertions.EqualValues(tt.want1, got1)
		})
	}
}
