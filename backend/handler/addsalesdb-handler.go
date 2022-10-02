package handler

import (
	"fmt"
	"github.com/beto-ouverney/go-affiliates/backend/config"
	salescontroller "github.com/beto-ouverney/go-affiliates/backend/internal/controllers/sales-controller"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

// BindFile is a struct for binding file
type BindFile struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// AddSalesDB is the handler for the route /addsalesdb, add sales to database
// @Summary      Add sales to database
// @Description  Add sales to database
// @Tags         sales, add, database
// @Produce      json
// @Success      200  {string} json "{"version":"0.0.1"}"
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router / [post]
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

	// Get file path
	d := fmt.Sprintf("%s%s", config.PATHFILE, file.Filename)

	if err := c.SaveUploadedFile(file, d); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf(`{"message":"%s"}`, err.Error()))
		return
	}

	// access sales controller
	ctl := salescontroller.New()
	res, err := ctl.Add(c, d)

	if err != nil {
		r, status := errorHandler(err)
		c.String(status, r)

		return
	}

	c.JSON(http.StatusOK, res)
	return
}
