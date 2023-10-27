package controllers

import (
	http "DanielDDHM/sender/src/services/http"

	"github.com/gin-gonic/gin"
)

// Returns the status of the application (health)
func CheckHealth(ctx *gin.Context) {

	ctx.JSON(200, http.Presenter(true, []string{"SUCCESS"}, map[string]interface{}{
		"running": true,
	}))
}
