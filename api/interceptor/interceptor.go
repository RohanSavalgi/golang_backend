package interceptor

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SendSuccessRes(c *gin.Context, data interface{}, statusCode int) {
	if statusCode == 0 {
		statusCode, _ = strconv.Atoi(os.Getenv("DEFAULT_HTTP_SUCCESS_CODE"))
	}
	response := CreateResponse(true, data, nil, "")
	c.AbortWithStatusJSON(statusCode, response)
}

func SendErrRes(c *gin.Context, err interface{}, errorMessage string, statusCode int) {
	if statusCode == 0 {
		statusCode, _ = strconv.Atoi(os.Getenv("DEFAULT_HTTP_ERROR_CODE"))
	}
	if errorMessage == "" {
		errorMessage = os.Getenv("DEFAULT_ERROR_MSG")
	}
	response := CreateResponse(false, nil, err, errorMessage)
	c.AbortWithStatusJSON(statusCode, response)
}