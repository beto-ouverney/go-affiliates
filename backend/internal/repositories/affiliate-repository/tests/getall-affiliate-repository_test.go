package tests

import (
	"context"
	"errors"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	mocksaffiliaterepository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/affiliate-repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_affiliateRepository_GetByName(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		name string
	}

	mockAffiliates := []entities.Affiliate{
		{
			ID:         1,
			Name:       "Alberto Paz",
			ProducerId: 1,
		},
		{
			ID:         2,
			Name:       "Beto Ouverney Paz",
			ProducerId: 2,
		},
	}
	tests := []struct {
		describe string
		want     *[]entities.Affiliate
		want1    *customerror.CustomError
		msg      string
		msg1     string
	}{
		{
			describe: "Should be able to get all affiliates of a content producer by name",
			want:     &mockAffiliates,
			want1:    nil,
			msg:      "The list os affliates must be the same",
			msg1:     "The error must be nil",
		},
		{
			describe: "Should be able return a nil list affiliates entity return and a custom error if affliates don`t exist in database",
			want:     nil,
			want1: customerror.NewError(customerror.ENOTFOUND, "Not found", "affiliate_repository.GetByName",
				errors.New("sql: no rows in result set")),
			msg:  "The list must be nil",
			msg1: "The error must be a custom error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			ctx := context.Background()
			m := new(mocksaffiliaterepository.IAffiliateRepository)
			m.On("GetAll", mock.AnythingOfType("*context.emptyCtx")).Return(tt.want, tt.want1)

			got, got1 := m.GetAll(ctx)
			assertions.Equal(tt.want, got, tt.msg)
			assertions.Equal(tt.want1, got1, tt.msg1)
		})
	}
}
