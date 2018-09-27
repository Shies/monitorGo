package net

import (
	"strings"
	"net/smtp"
)

const (
	_SMTP_USER   = "your email username"
	_SMTP_PASS   = "your email password"
	_SMTP_SERVER = "your email server"
)

// Smtp struct info.
type Mail struct {
	user   string
	pass   string
	server string
}

func NewMail() *Mail {
	mail := &Mail{
		user:   _SMTP_USER,
		pass:   _SMTP_PASS,
		server: _SMTP_SERVER,
	}

	return mail
}

func (m *Mail) Send(to, subject, body, mailtype string) error {
	var (
		content_type string
	)
	hp := strings.Split(m.server, ":")
	auth := smtp.PlainAuth("", m.user, m.pass, hp[0])
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + m.user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(m.server, auth, m.user, send_to, msg)
	return err
}
