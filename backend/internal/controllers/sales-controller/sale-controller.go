package sales_controller

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	salesusecase "github.com/beto-ouverney/go-affiliates/backend/internal/usecases/sales-usecase"
	"github.com/jmoiron/sqlx"
)

// ISaleController presents the interface for the sale controller
type ISaleController interface {
	Add(ctx context.Context, nameFile string) (*ResponseMsg, *customerror.CustomError)
	GetAll(ctx context.Context) (*[]entities.SaleResponse, *customerror.CustomError)
}

type saleController struct {
	useCase salesusecase.ISalesUseCase
}

// New creates a new sale controller
func New(db *sqlx.DB) ISaleController {
	return &saleController{
		salesusecase.New(db),
	}
}
