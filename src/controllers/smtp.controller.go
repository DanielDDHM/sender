package controllers

import (
	http "DanielDDHM/sender/src/services/http"
	smtp "DanielDDHM/sender/src/services/smtp"

	"github.com/gin-gonic/gin"
)

func SendSmtpMail(ctx *gin.Context) {
	to := ctx.Query("to")
	subject := ctx.Query("subject")
	message := ctx.Query("message")

	result := smtp.SendMail(
		to,
		subject,
		message,
	)

	if result == false {
		ctx.JSON(403, http.Presenter(
			result,
			[]string{"ERROR_TRYING_TO_SEND_AN_EMAIL_VIA_SMTP"},
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
