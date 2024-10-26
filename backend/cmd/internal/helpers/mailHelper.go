package helpers

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(address string, subject string, body string) error {
	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	smtpServer := os.Getenv("SMTP_SERVER")
	port := os.Getenv("SMTP_PORT")
	to := []string{address}

	auth := smtp.PlainAuth("", from, password, smtpServer)

	msg := []byte("From: " + from + "\n" +
		"To: " + address + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-version: 1.0;\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\n\n" +
		"<b>" + body + "</b>")

	err := smtp.SendMail(smtpServer+":"+port, auth, from, to, msg)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
