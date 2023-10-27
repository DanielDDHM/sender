package sms_services

import (
	"DanielDDHM/sender/src/utils"
	"log"

	"github.com/arcturial/clickatell-go"
)

// send an sms text via click a tell
func SendText(destination []string, message string) bool {
	if len(destination) <= 0 || len(message) <= 0 {
		return false
	}

	content := clickatell.Message{
		Destination: destination,
		Body:        message,
	}

	log.Println("KEY" + utils.Env.CLICKATELL_KEY)

	handler := clickatell.Rest(utils.Env.CLICKATELL_KEY, nil)

	_, err := handler.Send(content)

	if err != nil {
		log.Panicln(
			"[DanielDDHM-sender] An error ocurred while tryng to send an sms. Error: " + err.Error(),
		)
		return false
	}

	return true
}
