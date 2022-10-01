package sales_controller

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
)

// Add add a sale or more in database and returns a jsom with a success message
func (c *saleController) Add(ctx context.Context, nameFile string) (*string, *customerror.CustomError) {

	err := c.useCase.Add(ctx, nameFile)
	if err != nil {
		return nil, err
	}

	msg := `{ "message": "Sales added successfully" }`

	return &msg, nil
}
