package service

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

type EmailService struct {
	host string
	port int
	from string
}

func NewEmailService(host string, port int, from string) *EmailService {
	return &EmailService{host: host, port: port, from: from}
}

// SendPasswordResetEmail mengirim email berisi link reset password.
func (s *EmailService) SendPasswordResetEmail(toEmail, resetLink string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.from)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Reset Password - Job Tracker")

	body := fmt.Sprintf(
		"Halo,\n\nKami menerima permintaan reset password untuk akun ini.\n\nKlik link berikut untuk membuat password baru:\n%s\n\nLink ini berlaku selama 15 menit. Kalau kamu tidak meminta ini, abaikan email ini.",
		resetLink,
	)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(s.host, s.port, "", "")
	return d.DialAndSend(m)
}
