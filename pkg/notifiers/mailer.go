package notifiers

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"path"
	"strings"
	"time"

	"github.com/katta/jabfinder/pkg/models"
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

type Mailer struct {
	EMail
	SMTP
}

func getMessage(filters *models.Filters) string {
	msg := "Vaccines are available in the following centers :"
	if filters != nil {
		msg = fmt.Sprintf("Vaccines are available in the following centers for %d+ from %s", filters.Age, time.Now().Format("02-01-2006"))
		if filters.Date != "" {
			msg = fmt.Sprintf("Vaccines are available in the following centers for %d+ from %s", filters.Age, filters.Date)
		}
	}
	return msg
}

func (m *Mailer) SendMail(body string) {
	message := gomail.NewMessage()

	receivers := strings.Split(m.To, ",")

	for _, receiver := range receivers {
		message.SetHeader("From", m.From)
		message.SetHeader("To", receiver)
		message.SetHeader("Subject", m.Subject)
		message.SetBody("text/html", body)

		//fmt.Printf("Sending mailer with config: Email: %s, Password: %s \n", m.Email, m.Password)
		dialer := gomail.NewDialer(m.Host, m.Port, m.Email, m.Password)
		dialer.SSL = true

		log.Printf("Sending email to - %s", receiver)
		err := dialer.DialAndSend(message)
		if err != nil {
			fmt.Printf("Error sending email: %v", err)
		}
	}
}

func (m *Mailer) Notify(sessions []models.FlatSession, filters *models.Filters) {
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
		Message:  getMessage(filters),
		Sessions: sessions,
	})

	m.SendMail(body.String())
}
