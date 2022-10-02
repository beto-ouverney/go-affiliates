package handler

import (
	salescontroller "github.com/beto-ouverney/go-affiliates/backend/internal/controllers/sales-controller"
	"github.com/gin-gonic/gin"
)

// GetAllSales  is the handler for the route /sales, to catch all sales
func GetAllSales(c *gin.Context) {

	ctl := salescontroller.New()

	res, err := ctl.GetAll(c)
	if err != nil {
		r, status := errorHandler(err)
		c.String(status, r)
		return
	}

	c.JSON(200, res)
	return

}
