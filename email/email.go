package email

import (
	"log"
	"net/smtp"
)

const from = "onlecoop@gmail.com"
const pass = "@proyectoizi"

// Send -> Envía un correo electrónico a cierta direccion
func Send(to, subject, body, contentType string) {
	//to := "foobarbazz@mailinator.com"
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-version: 1.0;\n" +
		"Content-Type: " + contentType + "; charset=\"UTF-8\";\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	//log.Print("sent, visit http://foobarbazz.mailinator.com")
}
