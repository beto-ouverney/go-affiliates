package tests

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/controllers/sales-controller/mocks"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_saleController_Add(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		nameFile string
	}
	tests := []struct {
		describe string
		args     args
		want     []byte
		want1    *customerror.CustomError
	}{
		{
			describe: "Should be able add sales in database withou error",
			args: args{
				nameFile: "test.txt",
			},
			want:  []byte(`{"message": "Sales added successfully"}`),
			want1: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			ctx := context.Background()
			m := new(mocks.ISaleController)
			m.On("Add", mock.AnythingOfType("*context.emptyCtx"), tt.args.nameFile).Return(tt.want, tt.want1)

			got, got1 := m.Add(ctx, tt.args.nameFile)
			assertions.Equal(tt.want, got)
			assertions.Equal(tt.want1, got1)

		})
	}
}
