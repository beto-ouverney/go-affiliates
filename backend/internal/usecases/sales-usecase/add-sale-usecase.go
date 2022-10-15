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
func returnProducerID(p []entities.Producer, name string) (id int64) {

	for _, v := range p {
		if v.Name == name {
			return v.ID
		}
	}
	return
}

// returnProductID returns the product ID
func returnProductID(p []entities.Product, name string, idProducer int64) (id int64) {

	for _, v := range p {
		if v.Name == name && v.ProducerId == idProducer {
			return v.ID
		}
	}
	return
}

// getProdProducSales returns all  products  from database with ID,
// and sales of the content producers
func getProducersProductSales(ctx context.Context, u *salesUseCase, dEntry []parser.DataEntry, producersAll []entities.Producer) (*[]entities.Producer, *[]entities.Product,
	*[]entities.Sale, *customerror.CustomError) {

	//remove all duplicate producers
	cp := removeDuplicate(producersAll)

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

		if d.Type == 1 {
			s := entities.Sale{
				ProductId:  returnProductID(*allProducts, d.Product, d.ProducerId),
				ProducerId: dEntry[i].ProducerId,
				Value:      dEntry[i].Value,
				Commission: dEntry[i].Commission,
				Date:       dEntry[i].Date,
			}
			sales = append(sales, s)
		}
	}

	return allProducers, allProducts, &sales, nil
}

// getSalesAffiliates returns affiliates sales.
func getSalesAffiliates(ctx context.Context, u *salesUseCase, dataEntryType2, dataEntryType3 []parser.DataEntry, allCP *[]entities.Producer, allP *[]entities.Product, allAffD []entities.Affiliate, dataEntryAff []parser.DataEntry) (
	*[]entities.Sale, *customerror.CustomError) {

	for i, dA := range dataEntryAff {
		for _, dP := range dataEntryType3 {
			if dA.Date == dP.Date {
				id := returnProducerID(*allCP, dP.Seller)
				idProd := returnProductID(*allP, dP.Product, id)
				dataEntryAff[i].ProducerId = id
				dataEntryAff[i].ProductId = idProd
				allAffD[i].ProducerId = id
			}
		}
		for _, dP := range dataEntryType2 {
			if dA.Date == dP.Date {
				dataEntryAff[i].Value = dP.Value
			}
		}
	}

	// remove duplicate affiliates
	aff := removeDuplicate(allAffD)

	errorC := u.affiliateRepository.Add(ctx, aff)
	if errorC != nil {
		return nil, customerror.NewError(customerror.EINVALID, "Error", "sales_usecase.AddSale", errorC)
	}

	allAff, errorC := u.affiliateRepository.GetAll(ctx)
	if errorC != nil {
		return nil, customerror.NewError(customerror.EINVALID, "Error", "sales_usecase.AddSale", errorC)
	}

	var salesA []entities.Sale
	for _, d := range dataEntryAff {

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
	var dataEntryType3 []parser.DataEntry
	var dataEntryType2 []parser.DataEntry
	var dataEntryAffiliates []parser.DataEntry
	var dataEntryAll []parser.DataEntry
	var producersAll []entities.Producer
	var affiliatesAll []entities.Affiliate

	file, err := os.Open(nameFile)
	if err != nil {
		return customerror.NewError(customerror.EINVALID, "Error", "sales_usecase.AddSale", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		// jump a blank line
		if line == "" {
			continue
		}

		count++
		dataEntry, err := parser.ParseLine(line, count)
		if err != nil {
			return customerror.NewError(customerror.ECONFLICT, "Error", "sales_usecase.AddSale", err)
		}

		allLines = append(allLines, line)
		dataEntryAll = append(dataEntryAll, dataEntry)

		if line[0:1] == "1" || line[0:1] == "2" {

			dataEntryProducers = append(dataEntryProducers, dataEntry)
			cp := entities.Producer{
				Name: dataEntry.Seller,
			}
			producersAll = append(producersAll, cp)
			if line[0:1] == "2" {
				dataEntryType2 = append(dataEntryType2, dataEntry)
			}
		} else if line[0:1] == "3" {
			dataEntryType3 = append(dataEntryType3, dataEntry)
		} else {
			dataEntryAffiliates = append(dataEntryAffiliates, dataEntry)

			affiliatesAll = append(affiliatesAll, entities.Affiliate{
				Name: dataEntry.Seller,
			})
		}
	}

	file.Close()

	// get all producers, products and sales producers
	allCP, allP, salesP, errorC := getProducersProductSales(ctx, u, dataEntryProducers, producersAll)
	if errorC != nil {
		return errorC
	}

	// get affiliates sales
	salesAff, errorC := getSalesAffiliates(ctx, u, dataEntryProducers, dataEntryType2, allCP, allP, affiliatesAll, dataEntryAffiliates)
	if errorC != nil {
		return errorC
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	//write sales of producers in database
	var error1, error2 *customerror.CustomError
	go func() {
		error1 = writerDBProducers(ctx, u, *salesP, &wg)

	}()

	go func() {
		error2 = writerDBAffiliates(ctx, u, *salesAff, &wg)
	}()
	//write sales of affiliates in database
	wg.Wait()

	if error1 != nil {
		return customerror.NewError(customerror.EINVALID, "Error", "sales_usecase.AddSale", error1)
	}

	if error2 != nil {
		return customerror.NewError(customerror.EINVALID, "Error", "sales_usecase.AddSale", error2)
	}

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
