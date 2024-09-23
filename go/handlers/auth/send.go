package auth

import (
	"fmt"
	"net/smtp"
)

func send(mail string, body string) error {
	// Connect to the remote SMTP server.
	c, err := smtp.Dial("localhost:25")
	if err != nil {
		return err
	}

	// Set the sender and recipient first
	if err := c.Mail("m@seraferal.com"); err != nil {
		return err
	}
	if err := c.Rcpt(mail); err != nil {
		return err
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(wc, body)
	if err != nil {
		return err
	}
	err = wc.Close()
	if err != nil {
		return err
	}

	// Send the QUIT command and close the connection.
	return c.Quit()
}
