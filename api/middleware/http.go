package middleware

import (
	"net/http"
	"strings"

	"application/interceptor"
	"application/logger"

	"github.com/gin-gonic/gin"
)

type ContentType string

const (
	ContentTypeJSON ContentType = "application/json"
)

func VerifyContentType(t ContentType) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		if c.Request.Header.Get("Content-Type") != string(t) && strings.Split(c.Request.Header.Get("Content-Type"), ";")[0] != string(t) {
			logger.ThrowErrorLog("Invalid content type")
			interceptor.SendErrRes(c, "invalid content type", "Invalid content type", http.StatusBadRequest)
			return
		}
		c.Next()
	}
}
