package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Regex expression to verify if line of file is in correct format
const pattern = "(?P<type>[1-4])(?P<date>(19[0-9][0-9]|20[0-9][0-9])-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])(T)(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9]([\\+-])([01]\\d|2[0-3]):?([0-5]\\d)))(?P<product>[A-Z -]*)(?P<value>[0-9]*)(?P<seller>[A-Z ]*)"

// DataEntry presents a data line in the data file
// ID presents the line in the file
type DataEntry struct {
	ID         string
	Type       int
	Date       string
	Product    string
	ProductId  int64
	Value      int
	Commission int
	Seller     string
	ProducerId int64
}

// getMatchedValueByIdentifier returns the value of the matched expression by identifier
func getMatchedValueByIdentifier(identifier string, matches []string, expNames []string) string {

	for i, name := range expNames {
		if name == identifier {
			return matches[i]
		}
	}
	return "Incorrect format."
}

// ParseLine parser a string line and return a DataEntry struct
func ParseLine(line string, lineNumber int) (DataEntry, error) {
	errorLine := fmt.Sprintf("Line %d:", lineNumber)

	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(line)
	expNames := re.SubexpNames()

	lineOk := re.MatchString(line)
	if !lineOk {
		errorLine += " Incorrect format."
		return DataEntry{}, errors.New(errorLine)
	}

	typeString := getMatchedValueByIdentifier("type", matches, expNames)
	typeParam, _ := strconv.Atoi(typeString)
	date := getMatchedValueByIdentifier("date", matches, expNames)
	product := getMatchedValueByIdentifier("product", matches, expNames)
	if product == "" {
		errorLine += " Product is in incorrect format."
		return DataEntry{}, errors.New(errorLine)
	}
	if len(product) != 30 {
		errorLine += " Product must have length 30"
		return DataEntry{}, errors.New(errorLine)
	}
	product = strings.TrimSpace(product)
	valueString := getMatchedValueByIdentifier("value", matches, expNames)
	if len(valueString) != 10 {
		errorLine += "  Value must have 10 numbers."
		return DataEntry{}, errors.New(errorLine)
	}
	value, err := strconv.Atoi(valueString)
	if err != nil {
		errorLine += " Value is in incorrect format, must be a int number."
		return DataEntry{}, errors.New(errorLine)
	}
	seller := getMatchedValueByIdentifier("seller", matches, expNames)
	if seller == "" {
		errorLine += " Seller is in incorrect format."
		return DataEntry{}, errors.New(errorLine)
	}
	if len(seller) > 20 {
		errorLine += " Seller is too long, must be less than 21 characters."
		return DataEntry{}, errors.New(errorLine)
	}
	seller = strings.TrimSpace(seller)

	commission := 0
	if typeParam == 4 {
		commission = value
		value = 0
	}
	return DataEntry{
		ID:         line,
		Type:       typeParam,
		Date:       date,
		Product:    product,
		Commission: commission,
		Value:      value,
		Seller:     seller,
	}, nil

}
