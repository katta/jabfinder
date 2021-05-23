package mailer

import (
	"fmt"
	gomail "gopkg.in/mail.v2"
)

type EMail struct {
	From    string
	To      string
	Subject string
	Body    string
}

type SMTP struct {
	Host     string
	Port     int
	Email    string
	Password string
}

func SendMail(email EMail, smtpConfig SMTP) {
	message := gomail.NewMessage()

	message.SetHeader("From", email.From)
	message.SetHeader("To", email.To)
	message.SetHeader("Subject", email.Subject)
	message.SetBody("text/plain", email.Body)

	dialer := gomail.NewDialer(smtpConfig.Host, smtpConfig.Port, smtpConfig.Email, smtpConfig.Password)
	dialer.SSL = true

	err := dialer.DialAndSend(message)
	if err != nil {
		fmt.Printf("Error sending email: %v", err)
	}
}
