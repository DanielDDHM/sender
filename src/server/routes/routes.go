package routes

import (
	c "DanielDDHM/sender/src/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

// This function is responsible to create the application main route stack
func ConfigureApplicationRoutes(engine *gin.Engine) *gin.Engine {
	log.Print("[DanielDDHM-sender] Configuring all the application routes")

	main := engine.Group("api/v1")
	{
		health := main.Group("healthcheck")
		{
			health.GET("/", c.CheckHealth)
		}
		smtp := main.Group("smtp")
		{
			smtp.POST("/", c.SendSmtpMail)
		}
		sms := main.Group("sms")
		{
			sms.POST("/", c.SendSms)
		}
	}

	log.Print("[DanielDDHM-sender] All the routes was configured successfully")
	return engine
}
