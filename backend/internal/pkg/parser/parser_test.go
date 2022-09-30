package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseLineType(t *testing.T) {
	t.Log("Should be able parse the line type")
	dataEntry, err := ParseLine("12022-01-15T19:20:30-03:00CURSO DE BEM-ESTAR            0000012750JOSE CARLOS", 1)
	if err != nil {
		panic(err)
	}
	if dataEntry.Type != 1 {
		t.Errorf("Type should be 1, found %v", dataEntry.Type)
	}
}

func TestParseLineDate(t *testing.T) {
	t.Log("Should be able parse the line date")
	dataEntry, err := ParseLine("12022-01-15T19:20:30-03:00CURSO DE BEM-ESTAR            0000012750JOSE CARLOS", 1)
	if err != nil {
		panic(err)
	}
	if dataEntry.Date != "2022-01-15T19:20:30-03:00" {
		t.Errorf("Date should be 2022-01-15T19:20:30-03:00, found %v", dataEntry.Date)
	}

}

func TestParseLineProduct(t *testing.T) {
	t.Log("Should be able parse the line product")
	dataEntry, err := ParseLine("12022-01-15T19:20:30-03:00CURSO DE BEM-ESTAR            0000012750JOSE CARLOS", 1)
	if err != nil {
		panic(err)
	}
	if dataEntry.Product != "CURSO DE BEM-ESTAR" {
		t.Errorf("Product should be CURSO DE BEM-ESTAR, found %v", dataEntry.Product)
	}
}

func TestParseLineValue(t *testing.T) {
	t.Log("Should be able parse the line value")
	dataEntry, err := ParseLine("12022-01-15T19:20:30-03:00CURSO DE BEM-ESTAR            0000012750JOSE CARLOS", 1)
	if err != nil {
		panic(err)
	}
	if dataEntry.Value != 12750 {
		t.Errorf("Value should be 12750, found %v", dataEntry.Value)
	}
}

func TestParseLineSeller(t *testing.T) {
	t.Log("Should be able parse the line seller")
	dataEntry, err := ParseLine("12022-01-15T19:20:30-03:00CURSO DE BEM-ESTAR            0000012750JOSE CARLOS", 1)
	if err != nil {
		panic(err)
	}
	if dataEntry.Seller != "JOSE CARLOS" {
		t.Errorf("Seller should be JOSE CARLOS, found %v", dataEntry.Seller)
	}
}

func TestGetMatchedValueByIdentifier(t *testing.T) {
	t.Log("Should be able get the matched value by identifier")
	assertions := assert.New(t)

	t.Log("Should be able get the matched value by identifier")

	matches := []string{"TEST 1", "TEST 2", "TEST 3"}
	expNames := []string{"GROUP 1", "GROUP 2", "GROUP 3"}

	res := getMatchedValueByIdentifier("GROUP 1", matches, expNames)
	assertions.Equal("TEST 1", res, "Should be equal")

	res = getMatchedValueByIdentifier("GROUP 2", matches, expNames)
	assertions.Equal("TEST 2", res, "Should be equal")

	res = getMatchedValueByIdentifier("GROUP 3", matches, expNames)
	assertions.Equal("TEST 3", res, "Should be equal")

}

func TestVerifyErrorType(t *testing.T) {
	t.Log("Should be able verify the error type")
	assertions := assert.New(t)

	matches := []string{"A", "202-01-15T19:20:30-03:0", "TEST 3", "0000012750", "JOSE CARLOS"}
	expNames := []string{"type", "date", "product", "value", "seller"}

	err := verifyErrorType(matches, expNames)
	assertions.Equal(err, " Error in the type, must be a int number", "Should be return error")

	err = verifyErrorType(matches, expNames)
	assertions.NotNil(err, " Date, must be in format YYYY-MM-DDThh:mm:ss±hh:mm", "Should return error")
}

func TestParseLine(t *testing.T) {
	assertions := assert.New(t)
	t.Log("Should be able parse the line without error")

	line := "12022-01-15T19:20:30-03:00CURSO DE BEM-ESTAR            0000012750JOSE CARLOS"

	dataEntry := DataEntry{
		ID:      "12022-01-15T19:20:30-03:00CURSO DE BEM-ESTAR            0000012750JOSE CARLOS",
		Type:    1,
		Date:    "2022-01-15T19:20:30-03:00",
		Product: "CURSO DE BEM-ESTAR",
		Value:   12750,
		Seller:  "JOSE CARLOS",
	}

	parsedDataEntry, err := ParseLine(line, 1)

	assertions.Nil(err, "Should not return error")
	assertions.Equal(dataEntry, parsedDataEntry, "Should be equal")
}

func TestParseLineNameValue(t *testing.T) {
	assertions := assert.New(t)
	t.Log("Should be able parse the line and return name and value")

	line := "12022-01-15T19:20:30-03:00CURSO DE BEM-ESTAR            0000012750JOSE CARLOS"

	n, v := ParseLineNameValue(line)
	assertions.Equal(n, "JOSE CARLOS")
	assertions.Equal(v, 12750)
}
