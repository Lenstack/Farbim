package utils

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"math/rand"
	"time"
)

type IEmailManager interface {
	SendEmail(to string, subject string, body string) error
	GenerateCode() int64
}

type EmailManger struct {
	emailFrom    string
	passwordFrom string
}

func NewEmailManger(emailFrom string, passwordFrom string) *EmailManger {
	return &EmailManger{emailFrom: emailFrom, passwordFrom: passwordFrom}
}

func (em *EmailManger) SendEmail(to string, subject string, body string) error {
	message := gomail.NewMessage()
	message.SetHeader("To", to)
	message.SetHeader("From", em.emailFrom)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, em.emailFrom, em.passwordFrom)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := d.DialAndSend(message)
	if err != nil {
		return fmt.Errorf("cannot send an email: %v", err)
	}
	return nil
}

func (em *EmailManger) GenerateCode() int64 {
	rand.Seed(time.Now().UnixNano())
	max := 999999
	min := 100000
	code := min + rand.Intn(max-min)
	return int64(code)
}
