package sales_controller

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
)

// ResponseMsg is a struct that represents the response message to client Account model info
// @Description ResponseMsg response json message
type ResponseMsg struct {
	Message string `json:"message"`
}

// Add add a sale or more in database and returns a jsom with a success message
func (c *saleController) Add(ctx context.Context, nameFile string) (*ResponseMsg, *customerror.CustomError) {

	err := c.useCase.Add(ctx, nameFile)
	if err != nil {
		return nil, err
	}

	msg := ResponseMsg{
		Message: "Sales added successfully",
	}

	return &msg, nil
}
