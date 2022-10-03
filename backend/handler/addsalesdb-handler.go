package handler

import (
	"fmt"
	"github.com/beto-ouverney/go-affiliates/backend/config"
	salescontroller "github.com/beto-ouverney/go-affiliates/backend/internal/controllers/sales-controller"
	"github.com/beto-ouverney/go-affiliates/backend/internal/db"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"time"
)

// BindFile is a struct for binding file
type BindFile struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// AddSalesDB is the handler for the route /addsalesdb, add sales to database
// @Summary      Add sales to database
// @Description  Add sales from file to database
// @Tags         sales, add, database
// @Consumes     multipart/form-data
// @Param        multipart/form-data body BindFile true "File with sales"
// @Produce      json
// @Success      200  {object}  salescontroller.ResponseMsg
// @Failure      400  {object}  salescontroller.ResponseMsg
// @Failure      404  {object}  salescontroller.ResponseMsg
// @Failure      500  {object}  salescontroller.ResponseMsg
// @Router /sales/upload [post]
func AddSalesDB(c *gin.Context) {
	var bindFile BindFile

	// Bind file
	if err := c.ShouldBind(&bindFile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Field validation for 'File' failed on the 'required' tag. File not found"})
		return
	}

	// Save uploaded file
	file := bindFile.File

	if file.Header.Get("Content-Type") != "text/plain" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "File must be a text/plain"})
		return
	}

	conn := db.ConnectDB()
	defer conn.Close()

	// Get file path
	t := time.Now()
	d := fmt.Sprintf("%s%s-%s", config.PATHFILE, file.Filename, t)

	if err := c.SaveUploadedFile(file, d); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf(`{"message":"%s"}`, err.Error()))
		return
	}

	// access sales controller
	ctl := salescontroller.New(conn)
	res, err := ctl.Add(c.Request.Context(), d)

	if err != nil {
		r, status := errorHandler(err)
		c.String(status, r)

		return
	}

	c.JSON(http.StatusOK, res)
	return
}
