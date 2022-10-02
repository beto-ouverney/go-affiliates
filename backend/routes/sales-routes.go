package routes

import (
	"github.com/beto-ouverney/go-affiliates/backend/handler"
	"github.com/gin-gonic/gin"
)

// AddRoutes adds all sales routes to gin
func salesRoute(router *gin.RouterGroup) {
	r := router.Group("/sales")
	r.POST("/upload", handler.AddSalesDB)
	r.GET("/", handler.GetAllSales)
}
