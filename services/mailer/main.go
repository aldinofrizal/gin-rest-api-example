package mailer

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"gopkg.in/gomail.v2"
)

var Dialer *gomail.Dialer

func InitDialer() {
	Dialer = gomail.NewDialer(
		os.Getenv("MAIL_HOST"),
		587,
		os.Getenv("MAIL_USER"),
		os.Getenv("MAIL_PASSWORD"),
	)
}

type Mailer struct {
	To      []string
	Subject string
	Body    string
}

func (m *Mailer) SendMail() {
	mail := gomail.NewMessage()
	mail.SetHeader("From", "bumpa@heyho.com")
	mail.SetHeader("To", m.To...)
	mail.SetHeader("Subject", m.Subject)
	mail.SetBody("text/html", m.Body)

	if err := Dialer.DialAndSend(mail); err != nil {
		panic(err)
	} else {
		fmt.Printf("sent mail %s to %s", m.Subject, m.To)
	}
}

func (m *Mailer) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles("./services/mailer/" + templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	m.Body = buf.String()
	return nil
}
