package sales_controller

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
)

// GetAll returns all sales in database
func (c *saleController) GetAll(ctx context.Context) (*[]entities.SaleResponse, *customerror.CustomError) {

	all, err := c.useCase.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return all, nil
}
