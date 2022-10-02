package main

import (
	"github.com/beto-ouverney/go-affiliates/backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	gin.Logger()
	gin.Recovery()
	gin.SetMode(gin.ReleaseMode)
	gin.ForceConsoleColor()
	app := gin.New()

	router := app.Group("/api/v1")
	routes.AddRoutes(router)

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	log.Println("🚀 Server running on port: ", port)

	err = app.Run(port)
	if err != nil {
		log.Printf("Failed to launch api server:%+v\n", err)
	}

}
