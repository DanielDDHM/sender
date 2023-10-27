package controllers

import (
	http "DanielDDHM/sender/src/services/http"
	sms "DanielDDHM/sender/src/services/sms"

	"github.com/gin-gonic/gin"
)

func SendSms(ctx *gin.Context) {
	to := ctx.Query("to")
	message := ctx.Query("message")

	result := sms.SendText(
		[]string{to},
		message,
	)

	if result == false {
		ctx.JSON(403, http.Presenter(
			result,
			[]string{"ERROR_TRYING_TO_SEND_AN_SMS"},
			map[string]interface{}{},
		))
		return
	}

	ctx.JSON(200, http.Presenter(
		result,
		[]string{"SUCCESS"},
		map[string]interface{}{},
	))

}
