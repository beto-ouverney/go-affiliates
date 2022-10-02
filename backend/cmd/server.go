package main

import (
	_ "github.com/beto-ouverney/go-affiliates/backend/docs"
	"github.com/beto-ouverney/go-affiliates/backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

// @title           Go Affiliates API
// @version         1.0
// @description     This is server to manager content producers and affiliates sales.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    https://www.linkedin.com/in/beto-ouverney-paz/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8088
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	gin.Logger()
	gin.Recovery()
	gin.SetMode(gin.ReleaseMode)
	gin.ForceConsoleColor()
	app := gin.New()

	router := app.Group("/api/v1")
	routes.AddRoutes(router)

	app.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
