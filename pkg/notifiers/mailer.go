package notifiers

import (
	"bytes"
	"fmt"
	"github.com/katta/jabfinder/pkg/models"
	gomail "gopkg.in/mail.v2"
	"html/template"
	"log"
	"path"
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

type Mailer struct {
	EMail
	SMTP
}

func (m *Mailer) SendMail(body string) {
	message := gomail.NewMessage()

	message.SetHeader("From", m.From)
	message.SetHeader("To", m.To)
	message.SetHeader("Subject", m.Subject)
	message.SetBody("text/html", body)

	//fmt.Printf("Sending mailer with config: Email: %s, Password: %s \n", m.Email, m.Password)
	dialer := gomail.NewDialer(m.Host, m.Port, m.Email, m.Password)
	dialer.SSL = true

	err := dialer.DialAndSend(message)
	if err != nil {
		fmt.Printf("Error sending email: %v", err)
	}
}

func (m *Mailer) Notify(sessions []models.FlatSession) {
	//fmt.Printf("Sending mailer notification for sessions: %+v \n", sessions)

	var body bytes.Buffer

	emailTempl, err := template.ParseFiles(path.Join(".", "templates", "mail-notification.html"))
	if err != nil {
		log.Printf("Error parsing email template: %v", err)
		return
	}

	emailTempl.Execute(&body, struct {
		Message  string
		Sessions []models.FlatSession
	}{
		Message:  "Vaccines are available in the following centers :",
		Sessions: sessions,
	})

	m.SendMail(body.String())
}
