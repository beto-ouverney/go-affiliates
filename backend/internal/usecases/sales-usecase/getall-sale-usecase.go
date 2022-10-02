package sales_usecase

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
)

func (u *salesUseCase) GetAll(ctx context.Context) (*[]entities.SaleResponse, *customerror.CustomError) {

	sP, err := u.saleRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	sA, err := u.saleAffiliateRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	all := append(*sP, *sA...)

	return &all, nil
}
