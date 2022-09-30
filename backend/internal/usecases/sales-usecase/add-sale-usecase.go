package sales_usecase

import (
	"bufio"
	"context"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	"github.com/beto-ouverney/go-affiliates/backend/internal/pkg/parser"
	"os"
	"sync"
)

// returnProducerID returns the content producer ID
func returnProducerID(p []entities.Producer, name string) int64 {
	for _, v := range p {
		if v.Name == name {
			return v.ID
		}
	}
	return 0
}

// returnProductID returns the content product ID
func returnProductID(p []entities.Product, name string) int64 {
	for _, v := range p {
		if v.Name == name {
			return v.ID
		}
	}
	return 0
}

// getProdProducSales returns all producers from database with ID, alll products from database with ID and content producers sales
func getProducersProductSales(ctx context.Context, u *salesUseCase, dEntry []parser.DataEntry, cpAll []entities.Producer) (*[]entities.Producer,
	*[]entities.Product, *[]entities.Sale, *customerror.CustomError) {

	//remove all duplicate producers
	cp := removeDuplicate(cpAll)

	errorC := u.producerRepository.Add(ctx, cp)
	if errorC != nil {
		return nil, nil, nil, customerror.NewError(customerror.EINVALID, "Error", "sales_usecase.AddSale", errorC)
	}

	allProducers, errorC := u.producerRepository.GetAll(ctx)
	if errorC != nil {
		return nil, nil, nil, customerror.NewError(customerror.EINVALID, "Error", "sales_usecase.AddSale", errorC)
	}

	for i, d := range dEntry {

		for _, cp := range *allProducers {
			if cp.Name == d.Seller {
				dEntry[i].ProducerId = cp.ID
			}
		}
	}

	var pAll []entities.Product
	for _, d := range dEntry {
		pU := entities.Product{
			Name:       d.Product,
			ProducerId: d.ProducerId,
		}
		pAll = append(pAll, pU)
	}

	p := removeDuplicate(pAll)

	errorC = u.productRepository.Add(ctx, p)
	if errorC != nil {
		return nil, nil, nil, customerror.NewError(customerror.EINVALID, "Error", "sales_usecase.AddSale", errorC)
	}

	allProducts, errorC := u.productRepository.GetAll(ctx)
	if errorC != nil {
		return nil, nil, nil, customerror.NewError(customerror.EINVALID, "Error", "sales_usecase.AddSale", errorC)
	}

	var sales []entities.Sale

	for i, d := range dEntry {
		for _, p := range *allProducts {
			if p.Name == d.Product {
				dEntry[i].ProductId = p.ID
			}
		}

		s := entities.Sale{
			ProductId:  dEntry[i].ProductId,
			ProducerId: dEntry[i].ProducerId,
			Value:      dEntry[i].Value,
			Commission: dEntry[i].Commission,
			Date:       dEntry[i].Date,
		}
		sales = append(sales, s)
	}

	return allProducers, allProducts, &sales, nil
}

// getSalesAffiliates returns affiliates sales.
func getSalesAffiliates(ctx context.Context, u *salesUseCase, dataEntryProducers []parser.DataEntry, allLines []string, allCP *[]entities.Producer, allP *[]entities.Product) (
	*[]entities.Sale, *customerror.CustomError) {
	var aff []entities.Affiliate

	var dAff []parser.DataEntry
	for _, d := range dataEntryProducers {
		if d.Type == 2 {
			for _, l := range allLines {
				if l[0:1] == "4" {
					name, value := parser.ParseLineNameValue(l)
					idP := returnProducerID(*allCP, name)
					d := parser.DataEntry{
						Type:       4,
						Seller:     name,
						Value:      d.Value,
						ProducerId: idP,
						ProductId:  returnProductID(*allP, d.Product),
						Commission: value,
						Date:       d.Date,
					}
					dAff = append(dAff, d)
					affU := entities.Affiliate{
						Name:       name,
						ProducerId: idP,
					}
					aff = append(aff, affU)
				}
			}

		}
	}

	errorC := u.affiliateRepository.Add(ctx, aff)
	if errorC != nil {
		return nil, customerror.NewError(customerror.EINVALID, "Error", "sales_usecase.AddSale", errorC)
	}

	allAff, errorC := u.affiliateRepository.GetAll(ctx)
	if errorC != nil {
		return nil, customerror.NewError(customerror.EINVALID, "Error", "sales_usecase.AddSale", errorC)
	}

	var salesA []entities.Sale
	for _, d := range dAff {

		s := entities.Sale{
			ProductId:  d.ProductId,
			ProducerId: d.ProducerId,
			Value:      d.Value,
			Commission: d.Commission,
			Date:       d.Date,
		}
		for _, a := range *allAff {
			if a.Name == d.Seller {
				s.AffiliateId = a.ID
			}
		}
		salesA = append(salesA, s)
	}
	return &salesA, nil
}

// Add save the sales, the content producers, products and affiliates in the database. This function
// is responsible for reading the file, parsing the data, removing duplicates, saving the data in the database
func (u *salesUseCase) Add(ctx context.Context, nameFile string) *customerror.CustomError {
	var allLines []string
	var dataEntryProducers []parser.DataEntry
	var cpAll []entities.Producer

	path := "../tmp/"
	path += nameFile

	file, err := os.Open(path)
	if err != nil {
		return customerror.NewError(customerror.EINVALID, "Error", "sales_usecase.AddSale", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		count++
		dataEntry, err := parser.ParseLine(line, count)
		if err != nil {
			return customerror.NewError(customerror.ECONFLICT, "Error", "sales_usecase.AddSale", err)
		}

		allLines = append(allLines, line)

		if line[0:1] == "1" || line[0:1] == "2" {

			dataEntryProducers = append(dataEntryProducers, dataEntry)
			cp := entities.Producer{
				Name: dataEntry.Seller,
			}
			cpAll = append(cpAll, cp)
		}
	}

	allCP, allP, salesP, errorC := getProducersProductSales(ctx, u, dataEntryProducers, cpAll)
	if errorC != nil {
		return errorC
	}

	salesAff, errorC := getSalesAffiliates(ctx, u, dataEntryProducers, allLines, allCP, allP)
	if errorC != nil {
		return errorC
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	//write sales of producers in database

	errorC = writerDBProducers(ctx, u, *salesP, &wg)
	if errorC != nil {
		return customerror.NewError(customerror.EINVALID, "Error", "sales_usecase.AddSale", errorC)
	}

	//write sales of affiliates in database
	errorC = writerDBAffiliates(ctx, u, *salesAff, &wg)
	if errorC != nil {
		return customerror.NewError(customerror.EINVALID, "Error", "sales_usecase.AddSale", errorC)
	}
	wg.Wait()

	return nil
}

// writerDBProducers content producers write sales of the content producers in the database
func writerDBProducers(ctx context.Context, u *salesUseCase, sales []entities.Sale, wg *sync.WaitGroup) *customerror.CustomError {
	defer wg.Done()
	errorC := u.saleRepository.Add(ctx, sales)
	if errorC != nil {
		return customerror.NewError(customerror.EINVALID, "Error", "sales_usecase.AddSale", errorC)
	}
	return nil
}

// writerDBAffiliates content affiliates write sales of the content affiliates in the database
func writerDBAffiliates(ctx context.Context, u *salesUseCase, sales []entities.Sale, wg *sync.WaitGroup) *customerror.CustomError {
	defer wg.Done()

	errorC := u.saleAffiliateRepository.Add(ctx, sales)
	if errorC != nil {
		return customerror.NewError(customerror.EINVALID, "Error", "sales_usecase.AddSale", errorC)
	}
	return nil
}
