package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

// ConnectDB connects to the database
func ConnectDB() *sqlx.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("DB_CONNECTION"))
	return sqlx.MustConnect("postgres", os.Getenv("DB_CONNECTION"))
}
