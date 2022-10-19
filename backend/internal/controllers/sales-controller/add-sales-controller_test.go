package sales_controller

import (
	"context"
	"fmt"
	"github.com/beto-ouverney/go-affiliates/backend/config"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/usecases/sales-usecase/mocks"
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
		want     *ResponseMsg
		want1    *customerror.CustomError
	}{
		{
			describe: "Should be able add sales in database without error",
			args: args{
				nameFile: "test.txt",
			},
			want:  &ResponseMsg{Message: "Sales added successfully"},
			want1: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			d := fmt.Sprintf("%s%s", config.PATHFILE, tt.args.nameFile)
			ctx := context.Background()

			m := new(mocks.ISalesUseCase)

			m.On("Add", mock.AnythingOfType("*context.emptyCtx"), d).Return(tt.want1)

			c := saleController{m}
			got, got1 := c.Add(ctx, d)

			assertions.Equal(tt.want, got)
			assertions.Equal(tt.want1, got1)

		})
	}
}
