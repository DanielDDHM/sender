package http_services

import (
	"github.com/gin-gonic/gin"
)

// Returns an format to HTTP JSON requests
func Presenter(valid bool, messages []string, result map[string]interface{}) gin.H {
	return gin.H{
		"valid":    valid,
		"messages": messages,
		"result":   result,
	}
}
