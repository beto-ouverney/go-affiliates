package tests_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	salescontroller "github.com/beto-ouverney/go-affiliates/backend/internal/controllers/sales-controller"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	"github.com/beto-ouverney/go-affiliates/backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"mime/multipart"
	"net/http/httptest"
	"net/textproto"
	"os"
	"strings"
	"testing"
)

func TestNoSendFile(t *testing.T) {
	assertions := assert.New(t)
	t.Log("Return a error when no send file")

	app := gin.Default()
	gin.Logger()
	gin.Recovery()
	gin.SetMode(gin.DebugMode)

	router := app.Group("/api/v1")
	routes.AddRoutes(router)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/sales/upload", nil)

	app.ServeHTTP(res, req)

	t.Log(res.Body.String())
	msg := salescontroller.ResponseMsg{
		Message: `Field validation for 'File' failed on the 'required' tag. File not found`,
	}
	var actual salescontroller.ResponseMsg
	err := json.Unmarshal(res.Body.Bytes(), &actual)
	if err != nil {
		log.Fatal(err)
	}

	assertions.Equal(400, res.Code)
	assertions.Equal(msg, actual)

}

func TestAddSalesDB(t *testing.T) {
	assertions := assert.New(t)
	t.Setenv("POSTGRES_USER", "root")
	t.Setenv("POSTGRES_PASSWORD", "password")
	t.Setenv("POSTGRES_DB", "affiliates_db_test")
	t.Setenv("DB_CONNECTION", "postgres://root:password@affiliates_db_test:5432/affiliates_db_test?sslmode=disable")
	t.Log(os.Getenv("DB_CONNECTION"))

	//Verify if the database is in the test environment
	if strings.Contains(os.Getenv("POSTGRES_DB"), "test") {
		//defer dropDBTest(t)
	} else {
		t.Skip("Skipping test because it is not a test environment, the port number is not 6306")
	}

	app := gin.Default()
	gin.Logger()
	gin.Recovery()
	gin.SetMode(gin.DebugMode)

	router := app.Group("/api/v1")
	routes.AddRoutes(router)

	tests := []struct {
		describe             string
		args                 []string
		expectedStatus       int
		expectedBody         salescontroller.ResponseMsg
		msgStatus            string
		msgBody              string
		wantDBProducer       *[]entities.Producer
		wantDBAffiliate      *[]entities.Affiliate
		wantDBProducts       *[]entities.Product
		wantDBSaleProducer   *[]entities.Sale
		wantDBSalesAffiliate *[]entities.Sale
	}{
		{
			describe: "Should be able return a 200 status code and a success message if the all data was added to the database",
			args: []string{"12022-01-15T19:20:30-03:00CURSO GOLANG                  0000012750ALBERTO PAZ",
				"22022-01-16T14:13:54-03:00CURSO GOLANG                  0000012750JJ JAMESON",
				"32022-01-16T14:13:54-03:00CURSO GOLANG                  0000004500JJ JAMESON",
				"42022-01-16T14:13:54-03:00CURSO GOLANG                  0000004500ALBERTO PAZ",
			},
			expectedStatus: 200,
			expectedBody: salescontroller.ResponseMsg{
				Message: "Sales added successfully",
			},
			msgStatus: "The status code must be 200",
			msgBody:   "The body must be a success message",
			wantDBProducer: &[]entities.Producer{
				{ID: 1, Name: "ALBERTO PAZ"},
				{ID: 2, Name: "JJ JAMESON"},
			},
			wantDBAffiliate: &[]entities.Affiliate{
				{ID: 1, Name: "ALBERTO PAZ",
					ProducerId: 2},
			},
			wantDBProducts: &[]entities.Product{
				{
					ID:         1,
					Name:       "CURSO GOLANG",
					ProducerId: 1,
				},
				{
					ID:         2,
					Name:       "CURSO GOLANG",
					ProducerId: 2,
				},
			},
			wantDBSaleProducer: &[]entities.Sale{
				{
					ID:          1,
					ProductId:   1,
					ProducerId:  1,
					AffiliateId: 0,
					Value:       12750,
					Date:        "2022-01-15T19:20:30-03:00",
				},
			},
			wantDBSalesAffiliate: &[]entities.Sale{
				{
					ID:          1,
					ProducerId:  2,
					ProductId:   2,
					AffiliateId: 1,
					Value:       12750,
					Commission:  4500,
					Date:        "2022-01-16T14:13:54-03:00",
				},
			},
		},
		{
			describe: "Should be able return a 200 status code and a success message and the all data was added to the database if the file have a blank line",
			args: []string{"12022-01-15T19:20:30-03:00CURSO GOLANG                  0000012750ALBERTO PAZ",
				"22022-01-16T14:13:54-03:00CURSO GOLANG                  0000012750JJ JAMESON",
				"",
				"32022-01-16T14:13:54-03:00CURSO GOLANG                  0000004500JJ JAMESON",
				"42022-01-16T14:13:54-03:00CURSO GOLANG                  0000004500ALBERTO PAZ",
			},
			expectedStatus: 200,
			expectedBody: salescontroller.ResponseMsg{
				Message: "Sales added successfully",
			},
			msgStatus: "The status code must be 200",
			msgBody:   "The body must be a success message",
			wantDBProducer: &[]entities.Producer{
				{ID: 1, Name: "ALBERTO PAZ"},
				{ID: 2, Name: "JJ JAMESON"},
			},
			wantDBAffiliate: &[]entities.Affiliate{
				{ID: 1, Name: "ALBERTO PAZ",
					ProducerId: 2},
			},
			wantDBProducts: &[]entities.Product{
				{
					ID:         1,
					Name:       "CURSO GOLANG",
					ProducerId: 1,
				},
				{
					ID:         2,
					Name:       "CURSO GOLANG",
					ProducerId: 2,
				},
			},
			wantDBSaleProducer: &[]entities.Sale{
				{
					ID:          1,
					ProductId:   1,
					ProducerId:  1,
					AffiliateId: 0,
					Value:       12750,
					Date:        "2022-01-15T19:20:30-03:00",
				},
			},
			wantDBSalesAffiliate: &[]entities.Sale{
				{
					ID:          1,
					ProducerId:  2,
					ProductId:   2,
					AffiliateId: 1,
					Value:       12750,
					Commission:  4500,
					Date:        "2022-01-16T14:13:54-03:00",
				},
			},
		}, {
			describe: "Should not be able add the data to the database if the file have a invalid format",
			args: []string{"12022-01-15T19:20:30-03:00CURSO GOLANG                  0000012750ALBERTO PAZ",
				"22022-01-16T14:13:54-03:00CURSO GOLANG                  0000012750JJ JAMESON",
				"32022-01-13:54-03:00INVALID FORMAT           0000004500JJ JAMESON",
				"42022-01-16T14:13:54-03:00CURSO GOLANG                  0000004500ALBERTO PAZ",
			},
			expectedStatus: 400,
			expectedBody: salescontroller.ResponseMsg{
				Message: "Line 3: Incorrect format.",
			},
			msgStatus:            "The status code must be 400",
			msgBody:              "The body must be a failure message",
			wantDBProducer:       nil,
			wantDBAffiliate:      nil,
			wantDBProducts:       nil,
			wantDBSaleProducer:   nil,
			wantDBSalesAffiliate: nil,
		},
		{
			describe: "Should not be able add the data to the database if the product have a invalid format",
			args: []string{"12022-01-15T19:20:30-03:00                  0000012750ALBERTO PAZ",
				"22022-01-16T14:13:54-03:00CURSO GOLANG                  0000012750JJ JAMESON",
				"32022-01-16T14:13:54-03:00CURSO GOLANG                  0000004500JJ JAMESON",
				"42022-01-16T14:13:54-03:00CURSO GOLANG                  0000004500ALBERTO PAZ",
			},
			expectedStatus: 400,
			expectedBody: salescontroller.ResponseMsg{
				Message: "Line 1: Product must have length 30",
			},
			msgStatus:            "The status code must be 400",
			msgBody:              "The body must be a failure message",
			wantDBProducer:       nil,
			wantDBAffiliate:      nil,
			wantDBProducts:       nil,
			wantDBSaleProducer:   nil,
			wantDBSalesAffiliate: nil,
		},
		{
			describe: "Should not be able add the data to the database if the value have a invalid format",
			args: []string{
				"12022-01-16T14:13:54-03:00CURSO GOLANG                  0000A04500JJ JAMESON",
				"12022-01-16T14:13:54-03:00CURSO GOLANG                  0000004500ALBERTO PAZ",
			},
			expectedStatus: 400,
			expectedBody: salescontroller.ResponseMsg{
				Message: "Line 1:  Value must have 10 numbers.",
			},
			msgStatus:            "The status code must be 400",
			msgBody:              "The body must be a failure message",
			wantDBProducer:       nil,
			wantDBAffiliate:      nil,
			wantDBProducts:       nil,
			wantDBSaleProducer:   nil,
			wantDBSalesAffiliate: nil,
		},
		{
			describe: "Should not be able add the data to the database if the seller have a invalid format",
			args: []string{
				"12022-01-16T14:13:54-03:00CURSO GOLANG                  0000004500",
			},
			expectedStatus: 400,
			expectedBody: salescontroller.ResponseMsg{
				Message: "Line 1: Seller is in incorrect format.",
			},
			msgStatus:            "The status code must be 400",
			msgBody:              "The body must be a failure message",
			wantDBProducer:       nil,
			wantDBAffiliate:      nil,
			wantDBProducts:       nil,
			wantDBSaleProducer:   nil,
			wantDBSalesAffiliate: nil,
		}, {
			describe: "Should not be able add the data to the database if product have just spaces",
			args: []string{
				"22022-01-16T14:13:54-03:00                              0000012750THIAGO OLIVEIRA",
			},
			expectedStatus: 400,
			expectedBody: salescontroller.ResponseMsg{
				Message: "Line 1: Product is in incorrect format.",
			},
			msgStatus:            "The status code must be 400",
			msgBody:              "The body must be a failure message",
			wantDBProducer:       nil,
			wantDBAffiliate:      nil,
			wantDBProducts:       nil,
			wantDBSaleProducer:   nil,
			wantDBSalesAffiliate: nil,
		}, {
			describe: "Should not be able add the data to the database if name have just spaces",
			args: []string{
				"32022-01-16T14:13:54-03:00CURSO DE BEM-ESTAR            0000004500            ",
			},
			expectedStatus: 400,
			expectedBody: salescontroller.ResponseMsg{
				Message: "Line 1: Seller is in incorrect format.",
			},
			msgStatus:            "The status code must be 400",
			msgBody:              "The body must be a failure message",
			wantDBProducer:       nil,
			wantDBAffiliate:      nil,
			wantDBProducts:       nil,
			wantDBSaleProducer:   nil,
			wantDBSalesAffiliate: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			dropDBTest(t)
			initDBTest(t)
			// write test file
			f, err := os.Create("test.txt")

			if err != nil {
				log.Fatal(err)
			}

			for _, a := range tt.args {
				l := fmt.Sprintf("%s%s", a, "\n")
				_, errW := f.WriteString(l)
				if errW != nil {
					t.Fatal(errW)
				}
			}
			f.Close()

			body := &bytes.Buffer{}

			writer := multipart.NewWriter(body)
			h := make(textproto.MIMEHeader)
			h.Set("Content-Disposition", `form-data; name="file"; filename="test.txt"`)
			h.Set("Content-Type", "text/plain")
			fw, err := writer.CreatePart(h)
			if err != nil {
			}
			file, err := os.Open("test.txt")
			if err != nil {
				t.Fatal(err)
			}
			_, err = io.Copy(fw, file)
			if err != nil {
				t.Fatal(err)
			}
			// Close multipart writer.
			writer.Close()

			res := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/sales/upload", bytes.NewReader(body.Bytes()))
			req.Header.Set("Content-Type", writer.FormDataContentType())
			app.ServeHTTP(res, req)

			assertions.Equal(tt.expectedStatus, res.Code, tt.msgStatus)
			var actual salescontroller.ResponseMsg
			err = json.Unmarshal(res.Body.Bytes(), &actual)
			if err != nil {
				t.Errorf("Error unmarshalling response: %v", err)
			}

			assertions.EqualValues(tt.expectedBody.Message, actual.Message, tt.msgBody)

			// verify DATABASE

			conn, err := sqlx.Open("postgres", os.Getenv("DB_CONNECTION"))
			if err != nil {
				t.Fatal(err)
			}
			defer conn.Close()

			// verify producers
			var producers []entities.Producer
			err = conn.Select(&producers, "SELECT * FROM producers")
			if err != nil {
				t.Fatal(err)
			}

			if tt.wantDBProducer != nil {
				assertions.EqualValues(*tt.wantDBProducer, producers, "The producers must be equal")
			} else {
				assertions.Equal(0, len(producers), "Must be empty")
			}
			// verify affiliates
			var affiliates []entities.Affiliate
			err = conn.Select(&affiliates, "SELECT * FROM affiliates")
			if err != nil {
				t.Fatal(err)
			}
			if tt.wantDBAffiliate != nil {
				assertions.EqualValues(*tt.wantDBAffiliate, affiliates, "The affiliates must be equal")
			} else {
				assertions.Equal(0, len(affiliates), "Must be empty")
			}
			// verify products
			var products []entities.Product
			err = conn.Select(&products, "SELECT * FROM products")
			if err != nil {
				t.Fatal(err)
			}
			if tt.wantDBProducts != nil {
				assertions.EqualValues(*tt.wantDBProducts, products, "The products must be equal")
			} else {
				assertions.Equal(0, len(products), "Must be empty")
			}
			// verify sales_producers
			var salesP []entities.Sale
			err = conn.Select(&salesP, "SELECT * FROM sales_producers")
			if err != nil {
				t.Fatal(err)
			}
			if tt.wantDBSaleProducer != nil {
				assertions.EqualValues(*tt.wantDBSaleProducer, salesP, "The sales_producers must be equal")
			} else {
				assertions.Equal(0, len(salesP), "Must be empty")
			}
			// verify sales_affiliates
			var salesA []entities.Sale
			err = conn.Select(&salesA, "SELECT * FROM sales_affiliates")
			if err != nil {
				t.Fatal(err)
			}

			if tt.wantDBSalesAffiliate != nil {
				t.Log(salesA)
				t.Log(*tt.wantDBSalesAffiliate)
				assertions.EqualValues(*tt.wantDBSalesAffiliate, salesA, "The sales_affiliates must be equal")
			} else {
				assertions.Equal(0, len(salesA), "Must be empty")
			}

		})
	}
}
