package mailer

import (
	"fmt"
	"time"
)

func RegisterMail(email, verificationCode string) {
	m := Mailer{To: []string{email}, Subject: "Wawu verify account!"}
	d := struct {
		Email string
		Code  string
	}{email, verificationCode}
	if err := m.ParseTemplate("registermail.html", d); err == nil {
		m.SendMail()
	} else {
		fmt.Println(err)
	}
}

func RecoveryMail(err string, path string) {
	m := Mailer{
		To:      []string{"bumpa.heyho@gmail.com"},
		Subject: "Recovery Mailer ContentApp",
		Body:    fmt.Sprintf("error: %s\n path: %s\n time: %s", err, path, time.Now()),
	}
	m.SendMail()
	fmt.Println("recovery mailer send!")
}
