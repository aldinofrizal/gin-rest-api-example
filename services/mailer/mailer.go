package mailer

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/gomail.v2"
)

var dialer *gomail.Dialer

func InitDialer() {
	dialer = gomail.NewDialer(os.Getenv("MAIL_HOST"), 587, os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASSWORD"))
}

func RegisterMail(email, verificationCode string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "bumpa@heyho.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", fmt.Sprintf("please verify your email to http://localhost:8080/api/v1/users/verify?verificationCode=%s", verificationCode))

	if err := dialer.DialAndSend(m); err != nil {
		panic(err)
	} else {
		fmt.Println("success send registration mail to ", email)
	}
}

func RecoveryMail(err string, path string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "bumpa@heyho.com")
	m.SetHeader("To", "bumpa.heyho@gmail.com")
	m.SetHeader("Subject", "RECOVERY MAILER CONTENTAPP!")
	m.SetBody("text/html", fmt.Sprintf("error: %s\n path: %s\n time: %s", err, path, time.Now()))

	if err := dialer.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println("recovery mailer send!")
}
