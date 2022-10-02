package handler

import (
	salescontroller "github.com/beto-ouverney/go-affiliates/backend/internal/controllers/sales-controller"
	"github.com/gin-gonic/gin"
)

// GetAllSales  is the handler for the route /sales, to catch all sales
// @Summary      Get sales from database
// @Description  Get all sales of the content producers/affiliates from database
// @Produce      json
// @Success      200  {array}   entities.SaleResponse
// @Failure      400  {object}  salescontroller.ResponseMsg
// @Failure      404  {object}  salescontroller.ResponseMsg
// @Failure      500  {object}  salescontroller.ResponseMsg
// @Router /sales [get]
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
