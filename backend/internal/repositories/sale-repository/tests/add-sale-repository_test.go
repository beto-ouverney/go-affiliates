package tests

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	mocks_sale_repository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/sale-repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_saleRepository_Add(t *testing.T) {
	assertions := assert.New(t)

	mockSales := []entities.Sale{
		{
			ProducerId:  1,
			AffiliateId: 1,
			ProductId:   1,
			Value:       1000,
			Commission:  100,
			Date:        "22022-02-03T20:51:59-03:00",
		},
		{
			ProducerId:  2,
			AffiliateId: 2,
			ProductId:   2,
			Value:       2000,
			Commission:  0,
			Date:        "22022-02-03T20:51:59-03:00",
		},
	}

	mockOneSale := []entities.Sale{
		{
			ProducerId:  1,
			AffiliateId: 1,
			ProductId:   1,
			Value:       1000,
			Commission:  100,
			Date:        "22022-02-03T20:51:59-03:00",
		},
	}

	tests := []struct {
		describe string
		args     []entities.Sale
		want     *customerror.CustomError
		msg      string
	}{
		{
			describe: "Should be able add sales",
			args:     mockSales,
			want:     nil,
			msg:      "The error should be nil",
		},
		{
			describe: "Should be able add one sale",
			args:     mockOneSale,
			want:     nil,
			msg:      "The error should be nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			ctx := context.Background()
			m := new(mocks_sale_repository.ISaleRepository)
			m.On("Add", mock.AnythingOfType("*context.emptyCtx"), tt.args).Return(tt.want)

			got := m.Add(ctx, tt.args)
			assertions.Equal(tt.want, got, tt.msg)
		})
	}
}
