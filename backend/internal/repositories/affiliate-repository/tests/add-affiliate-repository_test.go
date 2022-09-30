package tests

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	mocksaffiliaterepository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/affiliate-repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_affiliateRepository_Add(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		data []entities.Affiliate
	}

	tests := []struct {
		describe string
		args     args
		want     *customerror.CustomError
		msg      string
	}{
		{
			describe: "Should be able add a producer and return a nil error",
			args: args{
				data: []entities.Affiliate{
					{
						Name:       "Hubla best company to contents producers",
						ProducerId: 1,
					},
					{
						Name:       "Hubla, company to contents producers",
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
			m := new(mocksaffiliaterepository.IAffiliateRepository)
			m.On("Add", mock.AnythingOfType("*context.emptyCtx"), tt.args.data).Return(tt.want)

			got := m.Add(ctx, tt.args.data)
			assertions.Equal(tt.want, got, tt.msg)

		})
	}
}
