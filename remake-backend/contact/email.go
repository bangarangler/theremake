package contact

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/smtp"

	"github.com/bangarangler/theremake/remake-backend/dotenvConfig"
)

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unknown from server")
		}
	}
	return nil, nil
}

func SendEmailFromContact(msg MessageFromContactForm) error {
	pass := dotenvConfig.EmailPW
	to := []string{dotenvConfig.EmailUser}
	from := msg.Email
	subject := msg.Subject
	body := msg.Message
	format := "Subject: " + subject + "\n" + body + "\n\n" + from + "\n\n" + msg.Name
	b := []byte(format)
	fmt.Println(to)
	fmt.Println(to[0])

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	conn, err := net.Dial("tcp", "smtp.gmail.com:587")
	if err != nil {
		println(err)
		return err
	}

	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		println(err)
		return err
	}

	tlsconfig := &tls.Config{
		ServerName: smtpHost,
	}

	if err = c.StartTLS(tlsconfig); err != nil {
		println(err)
		return err
	}

	auth := LoginAuth(to[0], pass)

	if err = c.Auth(auth); err != nil {
		println(err)
		return err
	}

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, b)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Email Sent!")
	return nil
}
