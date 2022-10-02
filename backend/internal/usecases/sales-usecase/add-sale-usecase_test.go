package sales_usecase

import (
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	"github.com/beto-ouverney/go-affiliates/backend/internal/pkg/parser"
	mocksaffiliaterepository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/affiliate-repository/mocks"
	mocksproducerrepository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/producer-repository/mocks"
	mocksproductrepository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/product-repository/mocks"
	mockssaleaffiliaterepository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/sale-affiliate-repository/mocks"
	mockssalerepository "github.com/beto-ouverney/go-affiliates/backend/internal/repositories/sale-producers-repository/mocks"
	mockssalesusecase "github.com/beto-ouverney/go-affiliates/backend/internal/usecases/sales-usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sync"
	"testing"
)

func Test_getProducersProductSales(t *testing.T) {
	assertions := assert.New(t)

	mRepo := new(mockssalerepository.ISaleRepository)
	sRepo := new(mockssaleaffiliaterepository.ISaleAffiliateRepository)

	pRepo := new(mocksproducerrepository.IProducerRepository)
	prodRepo := new(mocksproductrepository.IProductRepository)
	affRepo := new(mocksaffiliaterepository.IAffiliateRepository)

	type args struct {
		ctx      context.Context
		u        *salesUseCase
		dEntry   []parser.DataEntry
		allLines []string
		cpAll    []entities.Producer
		allP     *[]entities.Product
	}

	tests := []struct {
		describe string
		args     args
		want     *[]entities.Producer
		want1    *[]entities.Product
		want2    *[]entities.Sale
		want3    *customerror.CustomError
	}{
		{
			describe: "Should be able return the sales producers",
			args: args{
				ctx: context.Background(),
				u: &salesUseCase{
					mRepo, sRepo, pRepo, prodRepo, affRepo,
				},
				dEntry: []parser.DataEntry{
					{
						ID:         "12022-01-15T19:20:30-03:00CURSO DE BEM-ESTAR            0000012750JOSE CARLOS",
						Type:       1,
						Date:       "2022-01-15T19:20:30-03:00",
						Product:    "CURSO DE BEM-ESTAR",
						Value:      12750,
						Seller:     "JOSE CARLOS",
						Commission: 0,
						ProducerId: 1,
					},
				},
				allLines: []string{
					"12022-01-15T19:20:30-03:00CURSO DE BEM-ESTAR            0000012750JOSE CARLOS",
					"22022-01-16T14:13:54-03:00CURSO DE BEM-ESTAR            0000012750THIAGO OLIVEIRA",
					"32022-01-16T14:13:54-03:00CURSO DE BEM-ESTAR            0000004500THIAGO OLIVEIRA",
					"42022-01-16T14:13:54-03:00CURSO DE BEM-ESTAR            0000004500JOSE CARLOS",
				},
				cpAll: []entities.Producer{
					{
						ID:   1,
						Name: "JOSE CARLOS",
					},
					{
						ID:   2,
						Name: "THIAGO OLIVEIRA",
					},
				},
				allP: &[]entities.Product{
					{
						ID:   1,
						Name: "CURSO DE BEM-ESTAR",
					},
					{
						ID:   2,
						Name: "CURSO DE BEM-ESTAR",
					},
				},
			},
			want: &[]entities.Producer{
				{
					ID:   1,
					Name: "JOSE CARLOS",
				},
				{
					ID:   2,
					Name: "THIAGO OLIVEIRA",
				},
			},
			want1: &[]entities.Product{
				{
					ID:   1,
					Name: "CURSO DE BEM-ESTAR",
				},
				{
					ID:   2,
					Name: "CURSO DE BEM-ESTAR",
				},
			},
			want2: &[]entities.Sale{
				{
					ID:         1,
					ProductId:  2,
					Value:      12750,
					Commission: 0,
					ProducerId: 1,
					Date:       "2022-01-15T19:20:30-03:00",
				},
			},
			want3: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {

			m := new(IFunctionsMock)
			m.On("getProducersProductSales", mock.AnythingOfType("*context.emptyCtx"), tt.args.u, tt.args.dEntry, tt.args.cpAll).Return(tt.want, tt.want1, tt.want2, tt.want3)

			got, got1, got2, got3 := m.getProducersProductSales(tt.args.ctx, tt.args.u, tt.args.dEntry, tt.args.cpAll)
			assertions.Equal(got, tt.want)
			assertions.Equal(got1, tt.want1)
			assertions.Equal(got2, tt.want2)
			assertions.Equal(got3, tt.want3)

		})
	}
}

func Test_getSalesAffiliates(t *testing.T) {
	assertions := assert.New(t)

	mRepo := new(mockssalerepository.ISaleRepository)
	sRepo := new(mockssaleaffiliaterepository.ISaleAffiliateRepository)
	pRepo := new(mocksproducerrepository.IProducerRepository)
	prodRepo := new(mocksproductrepository.IProductRepository)
	affRepo := new(mocksaffiliaterepository.IAffiliateRepository)

	type args struct {
		ctx                context.Context
		u                  *salesUseCase
		dataEntryProducers []parser.DataEntry
		allLines           []string
		allCP              *[]entities.Producer
		allP               *[]entities.Product
	}

	tests := []struct {
		describe string
		args     args
		want     *[]entities.Sale
		want1    *customerror.CustomError
	}{
		{
			describe: "Should be able return the sales affiliates",
			args: args{
				ctx: context.Background(),
				u: &salesUseCase{
					mRepo, sRepo, pRepo, prodRepo, affRepo,
				},
				dataEntryProducers: []parser.DataEntry{
					{
						ID:         "12022-01-15T19:20:30-03:00CURSO DE BEM-ESTAR            0000012750JOSE CARLOS",
						Type:       1,
						Date:       "2022-01-15T19:20:30-03:00",
						Product:    "CURSO DE BEM-ESTAR",
						Value:      12750,
						Seller:     "JOSE CARLOS",
						Commission: 0,
						ProducerId: 1,
					},
				},
				allLines: []string{
					"12022-01-15T19:20:30-03:00CURSO DE BEM-ESTAR            0000012750JOSE CARLOS",
					"22022-01-16T14:13:54-03:00CURSO DE BEM-ESTAR            0000012750THIAGO OLIVEIRA",
					"32022-01-16T14:13:54-03:00CURSO DE BEM-ESTAR            0000004500THIAGO OLIVEIRA",
					"42022-01-16T14:13:54-03:00CURSO DE BEM-ESTAR            0000004500JOSE CARLOS",
				},
				allCP: &[]entities.Producer{
					{
						ID:   1,
						Name: "JOSE CARLOS",
					},
					{
						ID:   2,
						Name: "THIAGO OLIVEIRA",
					},
				},
				allP: &[]entities.Product{
					{
						ID:   1,
						Name: "CURSO DE BEM-ESTAR",
					},
					{
						ID:   2,
						Name: "CURSO DE BEM-ESTAR",
					},
				},
			},
			want:  nil,
			want1: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {

			m := new(IFunctionsMock)
			m.On("getSalesAffiliates", mock.AnythingOfType("*context.emptyCtx"), tt.args.u, tt.args.dataEntryProducers, tt.args.allLines, tt.args.allCP, tt.args.allP).Return(tt.want, tt.want1)
			got, got1 := m.getSalesAffiliates(tt.args.ctx, tt.args.u, tt.args.dataEntryProducers, tt.args.allLines, tt.args.allCP, tt.args.allP)
			assertions.Equal(tt.want, got)
			assertions.Equal(tt.want1, got1)
		})
	}
}

func Test_returnProductID(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		p      []entities.Product
		name   string
		idProd int64
	}
	tests := []struct {
		describe string
		args     args
		want     int64
		msg      string
	}{
		{
			describe: "Should be able return the product id",
			args: args{
				p: []entities.Product{
					{
						ID:         1,
						Name:       "Product 1",
						ProducerId: 3,
					},
					{
						ID:         2,
						Name:       "Product 2",
						ProducerId: 3,
					},
				},
				name:   "Product 2",
				idProd: 3,
			},
			want: 2,
			msg:  "The product id must be 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			got := returnProductID(tt.args.p, tt.args.name, tt.args.idProd)
			assertions.Equal(tt.want, got, tt.msg)
		})
	}
}

func Test_salesUseCase_Add(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		nameFile string
	}
	tests := []struct {
		describe string
		args     args
		want     *customerror.CustomError
		msg      string
	}{
		{
			describe: "Shoulbe be able to add sales and return nil error",
			args: args{
				nameFile: "test1.txt",
			},
			want: nil,
			msg:  "Must be nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			ctx := context.Background()
			m := new(mockssalesusecase.ISalesUseCase)
			m.On("Add", mock.AnythingOfType("*context.emptyCtx"), tt.args.nameFile).Return(tt.want)

			got := m.Add(ctx, tt.args.nameFile)
			assertions.Equal(tt.want, got, tt.msg)
		})
	}
}

func Test_writerDBAffiliates(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		ctx   context.Context
		u     *salesUseCase
		sales []entities.Sale
		wg    *sync.WaitGroup
	}

	mRepo := new(mockssalerepository.ISaleRepository)
	sRepo := new(mockssaleaffiliaterepository.ISaleAffiliateRepository)
	pRepo := new(mocksproducerrepository.IProducerRepository)
	prodRepo := new(mocksproductrepository.IProductRepository)
	affRepo := new(mocksaffiliaterepository.IAffiliateRepository)

	tests := []struct {
		describe string
		args     args
		want     *customerror.CustomError
		msg      string
	}{
		{
			describe: "Should be able to write in DB and return nil error",
			args: args{
				ctx: context.Background(),
				u:   &salesUseCase{mRepo, sRepo, pRepo, prodRepo, affRepo},
				sales: []entities.Sale{
					{
						ProducerId:  1,
						ProductId:   1,
						AffiliateId: 1,
						Value:       10,
						Commission:  10,
						Date:        "2022-02-04T07:42:12-03:00",
					},
				},
				wg: &sync.WaitGroup{},
			},
			want: nil,
			msg:  "Must be nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {

			m := new(IFunctionsMock)
			m.On("writerDBAffiliates", mock.AnythingOfType("*context.emptyCtx"), tt.args.u, tt.args.sales,
				tt.args.wg).Return(tt.want)

			got := m.writerDBAffiliates(tt.args.ctx, tt.args.u, tt.args.sales, tt.args.wg)
			assertions.Equal(tt.want, got, tt.msg)
		})
	}
}

func Test_writerDBProducers(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		ctx   context.Context
		u     *salesUseCase
		sales []entities.Sale
		wg    *sync.WaitGroup
	}

	mRepo := new(mockssalerepository.ISaleRepository)
	sRepo := new(mockssaleaffiliaterepository.ISaleAffiliateRepository)
	pRepo := new(mocksproducerrepository.IProducerRepository)
	prodRepo := new(mocksproductrepository.IProductRepository)
	affRepo := new(mocksaffiliaterepository.IAffiliateRepository)

	tests := []struct {
		describe string
		args     args
		want     *customerror.CustomError
		msg      string
	}{
		{
			describe: "Should be able to write DB and return nil error",
			args: args{
				ctx: context.Background(),
				u:   &salesUseCase{mRepo, sRepo, pRepo, prodRepo, affRepo},
				sales: []entities.Sale{
					{
						ProducerId: 1,
						ProductId:  1,
						Value:      10,
						Date:       "2022-02-04T07:42:12-03:00",
					},
				},
				wg: &sync.WaitGroup{},
			},
			want: nil,
			msg:  "Must be nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {

			m := new(IFunctionsMock)
			m.On("writerDBProducers", mock.AnythingOfType("*context.emptyCtx"), tt.args.u, tt.args.sales,
				tt.args.wg).Return(tt.want)

			got := m.writerDBProducers(tt.args.ctx, tt.args.u, tt.args.sales, tt.args.wg)
			assertions.Equal(tt.want, got, tt.msg)
		})
	}
}
