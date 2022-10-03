package sales_usecase

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	affiliaterepository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/affiliate-repository"
	producerrepository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/producer-repository"
	productrepository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/product-repository"
	saleaffiliaterepository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/sale-affiliate-repository"
	salerepository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/sale-producers-repository"
	"github.com/jmoiron/sqlx"
)

// ISalesUseCase presents the interface for the sales use case
type ISalesUseCase interface {
	Add(ctx context.Context, nameFile string) *customerror.CustomError
	GetAll(ctx context.Context) (*[]entities.SaleResponse, *customerror.CustomError)
}

type salesUseCase struct {
	saleRepository          salerepository.ISaleRepository
	saleAffiliateRepository saleaffiliaterepository.ISaleAffiliateRepository
	producerRepository      producerrepository.IProducerRepository
	productRepository       productrepository.IProductRepository
	affiliateRepository     affiliaterepository.IAffiliateRepository
}

// New creates a new sales use case
func New(db *sqlx.DB) ISalesUseCase {
	return &salesUseCase{
		salerepository.New(db),
		saleaffiliaterepository.New(db),
		producerrepository.New(db),
		productrepository.New(db),
		affiliaterepository.New(db),
	}
}
