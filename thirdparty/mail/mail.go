package mail

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func SendSmtpMail(subject, to, message string) error {
	host := viper.GetString("mail_host")
	port := viper.GetInt("mail_port")
	uname := viper.GetString("mail_username")
	pass := viper.GetString("mail_password")

	mail := gomail.NewMessage()
	mail.SetHeader("From", viper.GetString("mail_sender"))
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", message)

	dial := gomail.NewDialer(host, port, uname, pass)

	if err := dial.DialAndSend(mail); err != nil {
		return err
	}

	return nil
}
