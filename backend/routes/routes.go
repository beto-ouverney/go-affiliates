package routes

import "github.com/gin-gonic/gin"

// AddRoutes adds all routes to gin
func AddRoutes(superRoute *gin.RouterGroup) {
	salesRoute(superRoute)
}
