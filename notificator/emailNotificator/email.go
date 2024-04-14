package emailnotificator

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

type EmailNotificator struct {
	Settings EmailSettings
}

type EmailSettings struct {
	From   string
	To     []string
	Cc     []string
	Subj   string
	Dialer gomail.Dialer
}

func (n EmailNotificator) SendAlert(a ...any) error {
	msg := fmt.Sprint(a...)
	err := n.sendEmail(n.Settings.To, n.Settings.Cc, n.Settings.Subj, msg)
	if err != nil {
		return err
	}
	return nil
}

func (n EmailNotificator) Create() EmailNotificator {
	return EmailNotificator{}
}

func (n *EmailNotificator) SetDialer(smtp string, port int, username, password string) {
	n.Settings.Dialer = *gomail.NewDialer(smtp, port, username, password)
}

func (n *EmailNotificator) SetSettings(from string, to, cc []string, subj string) {
	n.Settings = EmailSettings{
		From: from,
		To:   to,
		Cc:   cc,
		Subj: subj,
	}
}

func (n EmailNotificator) sendEmail(to, cc []string, subject string, body string, attachments ...string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", n.Settings.From)
	m.SetHeader("To", to...)
	m.SetHeader("Cc", cc...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	for _, attachment := range attachments {
		m.Attach(attachment)
	}

	if err := n.Settings.Dialer.DialAndSend(m); err != nil {
		return err
	} else {
		return nil
	}
}
