package sales_controller

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	salesusecase "github.com/beto-ouverney/go-affiliates/backend/internal/usecases/sales-usecase"
)

// ISaleController presents the interface for the sale controller
type ISaleController interface {
	Add(ctx context.Context, nameFile string) (*string, *customerror.CustomError)
}

type saleController struct {
	useCase salesusecase.ISalesUseCase
}

// New creates a new sale controller
func New() ISaleController {
	return &saleController{
		salesusecase.New(),
	}
}
