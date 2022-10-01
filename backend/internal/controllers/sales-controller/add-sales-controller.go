package sales_controller

import (
	"context"
	"encoding/json"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
)

// Add add a sale or more in database and returns a jsom with a success message
func (c *saleController) Add(ctx context.Context, nameFile string) ([]byte, *customerror.CustomError) {

	err := c.useCase.Add(ctx, nameFile)
	if err != nil {
		return nil, err
	}

	msg := struct {
		Message string `json:"message"`
	}{
		Message: "Sales added successfully",
	}

	msgJ, errJ := json.MarshalIndent(&msg, "", "  ")
	if errJ != nil {
		return nil, customerror.NewError(customerror.EINTERNAL, "Error", "sale_controller.Add", errJ)
	}

	return msgJ, nil
}
