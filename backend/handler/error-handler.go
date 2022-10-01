package handler

import (
	"fmt"
	"github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	"log"
)

// errorHandler handles the error and return the message and status code
func errorHandler(err *customerror.CustomError) (response string, status int) {

	if err.Code == customerror.ENOTFOUND {
		status = 404
	} else if err.Code == customerror.ECONFLICT {
		status = 400
	} else {
		log.Println(err)
		status = 500
	}
	response = fmt.Sprintf(`{"message":"%s"}`, err.Error())

	return
}
