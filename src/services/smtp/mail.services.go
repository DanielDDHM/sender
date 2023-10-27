package smtp_services

import (
	"DanielDDHM/sender/src/utils"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Send an email via sendgrid
func SendMail(toEmail string, subject string, content string) bool {
	from := mail.NewEmail("DanielDDHM", utils.Env.SMTP_HOST)
	to := mail.NewEmail(toEmail, toEmail)

	message := mail.NewSingleEmail(from, subject, to, content, content)

	client := sendgrid.NewSendClient(utils.Env.SMTP_PASSWORD)

	_, err := client.Send(message)

	if err != nil {
		log.Println("[DanielDDHM-sender] An error ocurred while tryng to send an email")
		log.Println(err)
		return false
	}

	log.Println("[DanielDDHM-sender] Email sent succesfully")

	return true
}
