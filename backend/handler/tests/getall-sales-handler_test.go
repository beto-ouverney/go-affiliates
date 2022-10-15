package tests_test

import (
	"encoding/json"
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	"github.com/beto-ouverney/go-affiliates/backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestGetAllSales(t *testing.T) {
	assertions := assert.New(t)
	t.Setenv("POSTGRES_USER", "root")
	t.Setenv("POSTGRES_PASSWORD", "password")
	t.Setenv("POSTGRES_DB", "affiliates_db_test")
	t.Setenv("DB_CONNECTION", "postgres://root:password@affiliates_db_test:5432/affiliates_db_test?sslmode=disable")
	t.Log(os.Getenv("DB_CONNECTION"))

	//Verify if the database is in the test environment
	if strings.Contains(os.Getenv("POSTGRES_DB"), "test") {
		initDBTest(t)
		defer dropDBTest(t)
	} else {
		t.Skip("Skipping test because it is not a test environment, the port number is not 6306")
	}

	app := gin.Default()
	gin.Logger()
	gin.Recovery()
	gin.SetMode(gin.DebugMode)

	router := app.Group("/api/v1")
	routes.AddRoutes(router)

	mockSales := []entities.SaleResponse{
		{
			Producer:   "ALBERTO PAZ",
			Product:    "CURSO GOLANG",
			Value:      1000,
			Commission: 0,
			Date:       "22022-02-03T20:51:59-03:00",
		},
		{
			Producer:   "ALBERTO PAZ",
			Affiliate:  "GEORGE MARTIN",
			Product:    "CURSO GOLANG",
			Value:      1000,
			Commission: 100,
			Date:       "22022-03-03T20:51:59-03:00",
		},
	}

	mockSalesJson, err := json.Marshal(mockSales)
	if err != nil {
		t.Fatal(err)
	}

	test := struct {
		describe       string
		expectedStatus int
		expectedBody   []byte
		msgStatus      string
		msgBody        string
	}{
		describe:       "Should be able to return all sales",
		expectedStatus: 200,
		expectedBody:   mockSalesJson,
		msgStatus:      "The status code is not the expected",
		msgBody:        "The body is not the expected",
	}

	//INSERT DATA IN DB TESTS

	schemaPopulateDB := [5]string{
		`INSERT INTO producers (name) VALUES ('ALBERTO PAZ');`,
		`INSERT INTO affiliates (name,producer_id) VALUES ('GEORGE MARTIN',1);`,
		`INSERT INTO products (name,producer_id) VALUES ('CURSO GOLANG',1);`,
		`INSERT INTO sales_producers (producer_id,product_id,value, date) VALUES (1,1,1000,'22022-02-03T20:51:59-03:00');`,
		`INSERT INTO sales_affiliates (producer_id, affiliate_id, product_id, value, commission, date) VALUES (1, 1, 1, 1000, 100, '22022-03-03T20:51:59-03:00');`,
	}

	conn, err := sqlx.Open("postgres", os.Getenv("DB_CONNECTION"))
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, s := range schemaPopulateDB {
		_, err = conn.Exec(s)
		if err != nil {
			t.Fatal(s)
			t.Fatal(err)
		}
	}

	t.Run(test.describe, func(t *testing.T) {
		res := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/api/v1/sales/", nil)
		app.ServeHTTP(res, req)

		assertions.Equal(test.expectedStatus, res.Code, test.msgStatus)

		assertions.Equal(test.expectedBody, res.Body.Bytes(), test.msgBody)
	})
}

func TestGetEmptyArrayIfNotHaveSales(t *testing.T) {
	assertions := assert.New(t)
	t.Setenv("POSTGRES_USER", "root")
	t.Setenv("POSTGRES_PASSWORD", "password")
	t.Setenv("POSTGRES_DB", "affiliates_db_test")
	t.Setenv("DB_CONNECTION", "postgres://root:password@affiliates_db_test:5432/affiliates_db_test?sslmode=disable")
	t.Log(os.Getenv("DB_CONNECTION"))

	//Verify if the database is in the test environment
	if strings.Contains(os.Getenv("POSTGRES_DB"), "test") {
		initDBTest(t)
		// when the test is finished, recreate db
		defer initDBTest(t)
	} else {
		t.Skip("Skipping test because it is not a test environment, the port number is not 6306")
	}

	app := gin.Default()
	gin.Logger()
	gin.Recovery()
	gin.SetMode(gin.DebugMode)

	router := app.Group("/api/v1")
	routes.AddRoutes(router)

	test := struct {
		describe       string
		expectedStatus int
		expectedBody   []byte
		msgStatus      string
		msgBody        string
	}{
		describe:       "Should be able to return a empty array if not have sales",
		expectedStatus: 200,
		expectedBody:   []byte(`[]`),
		msgStatus:      "The status code is not the expected",
		msgBody:        "The body is not the expected",
	}

	t.Run(test.describe, func(t *testing.T) {
		res := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/api/v1/sales/", nil)
		app.ServeHTTP(res, req)

		assertions.Equal(test.expectedStatus, res.Code, test.msgStatus)

		assertions.Equal(test.expectedBody, res.Body.Bytes(), test.msgBody)
	})
}
