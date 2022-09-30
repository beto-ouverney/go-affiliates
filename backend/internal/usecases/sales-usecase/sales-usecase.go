package sales_usecase

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	affiliaterepository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/affiliate-repository"
	producerrepository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/producer-repository"
	productrepository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/product-repository"
	saleaffiliaterepository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/sale-affiliate-repository"
	salerepository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/sale-producers-repository"
)

// ISalesUseCase presents the interface for the sales use case
type ISalesUseCase interface {
	Add(ctx context.Context, nameFile string) *customerror.CustomError
}

type salesUseCase struct {
	saleRepository          salerepository.ISaleRepository
	saleAffiliateRepository saleaffiliaterepository.ISaleAffiliateRepository
	producerRepository      producerrepository.IProducerRepository
	productRepository       productrepository.IProductRepository
	affiliateRepository     affiliaterepository.IAffiliateRepository
}

// New creates a new sales use case
func New() ISalesUseCase {
	return &salesUseCase{
		salerepository.New(),
		saleaffiliaterepository.New(),
		producerrepository.New(),
		productrepository.New(),
		affiliaterepository.New(),
	}
}
