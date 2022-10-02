package tests_test

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"testing"
)

var schemasInit = [5]string{
	`CREATE TABLE IF NOT EXISTS producers(
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    UNIQUE (name)
);`,
	`CREATE TABLE IF NOT EXISTS affiliates(
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    producer_id INT NOT NULL,
    FOREIGN KEY (producer_id) REFERENCES producers(id),
    UNIQUE (name, producer_id)
);`,
	`CREATE TABLE IF NOT EXISTS products(
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    producer_id INT NOT NULL,
    FOREIGN KEY (producer_id) REFERENCES producers(id),
    UNIQUE (name, producer_id)
);`,
	`CREATE TABLE IF NOT EXISTS sales_producers(
    id SERIAL PRIMARY KEY,
    producer_id INT NOT NULL,
    product_id INT NOT NULL,
    FOREIGN KEY (producer_id) REFERENCES producers(id),
    FOREIGN KEY (product_id) REFERENCES products(id),
    value INT NOT NULL,
    commission INT,
    date TIMESTAMPTZ,
    UNIQUE (producer_id, product_id, date)
);`,
	`CREATE TABLE IF NOT EXISTS sales_affiliates(
    id SERIAL PRIMARY KEY,
    producer_id INT NOT NULL,
    affiliate_id INT NOT NULL,
    product_id INT NOT NULL,
    FOREIGN KEY (producer_id) REFERENCES producers(id),
    FOREIGN KEY (affiliate_id) REFERENCES affiliates(id),
    FOREIGN KEY (product_id) REFERENCES products(id),
    value INT NOT NULL,
    commission INT,
    date TIMESTAMPTZ,
    UNIQUE (producer_id, product_id, affiliate_id, date)
);`,
}

var schemasDrop = [5]string{
	`DROP TABLE IF EXISTS sales_affiliates;`,
	`DROP TABLE IF EXISTS sales_producers;`,
	`DROP TABLE IF EXISTS products;`,
	`DROP TABLE IF EXISTS affiliates;`,
	`DROP TABLE IF EXISTS producers;`,
}

// DB_CONNECTION=postgres://root:password@affiliates_db_test:5432/affiliates_db_test?sslmode=disable
const POSTGREES_CONNECTION = "user=root password=password dbname=affiliates_db_test sslmode=disable"
const DATABASE_URL = "postgres://root:password@affiliates_db_test:5432/affiliates_db_test?sslmode=disable"

func initDBTest(t *testing.T) {
	t.Log("Initializing database test")
	t.Log(POSTGREES_CONNECTION)
	conn, err := sqlx.Open("postgres", DATABASE_URL)

	if err != nil {
		t.Fatal(err.Error())
	}

	for _, s := range schemasDrop {
		_, err = conn.Exec(s)
		if err != nil {
			t.Fatal(err)
		}
	}

	t.Log("Database test dropped successfully")

	t.Setenv("POSTGRES_DB", "affiliates_db_test")

	for _, s := range schemasInit {
		_, err = conn.Exec(s)
		if err != nil {
			t.Fatal(err)
		}
	}
	conn.Close()
	t.Log("Database test created successfully")
}

func dropDBTest(t *testing.T) {
	t.Log("Dropping database test")
	conn, err := sqlx.Open("postgres", DATABASE_URL)

	if err != nil {
		t.Fatal(err.Error())
	}

	for _, s := range schemasDrop {
		_, err = conn.Exec(s)
		if err != nil {
			t.Fatal(err)
		}
	}
	conn.Close()
	t.Log("Database test dropped successfully")
}
