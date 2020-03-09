package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	// grab our arguments. keeping the program name
	// to make it simpler to work with the arguments, since
	// the index will match the semantic order that way
	arguments := os.Args[:]

	// we need 7 arguments. they can't be empty, either.
	// 1) API Key
	// 2) Sender Name
	// 3) Sender Email
	// 4) Recipient Name
	// 5) Recipient Email
	// 6) Subject
	// 7) Plain Message Body
	if len(arguments) < 1 {
		log.Fatalf("\nERROR: This binary requires exactly 7 arguments:\n  (1) Sendgrid API Key\n  (2) Sender Name\n  (3) Sender Email\n  (4) Recipient Name\n  (5) Recipient Email\n  (6) Subject\n  (7) Message\n You have a missing or empty value in invocation of zabbix_sendgrid_media::main")
	}

	// now build out our message...
	sender := mail.NewEmail(arguments[2], arguments[3])
	recipient := mail.NewEmail(arguments[4], arguments[5])
	message := mail.NewSingleEmail(sender, arguments[6], recipient, arguments[7], arguments[7])

	// create the client, and send
	client := sendgrid.NewSendClient(arguments[1])
	smtpResponse, err := client.Send(message)
	// log if we error...
	if err != nil {
		log.Printf("\nERROR: sending email to %s <%s> from %s <%s> failed!\nSUBJECT: %s\nMESSAGE: %s\nERROR MESSAGE: %s\n", arguments[2], arguments[3], arguments[4], arguments[5], arguments[6], arguments[7], err)
		os.Exit(1)
	}

	// ...and log if we don't, but exit with 0.
	fmt.Printf("\nMessage sent to %s\nStatus: %d\nHeaders: %s,\bBody: %s\n", arguments[2], smtpResponse.StatusCode, smtpResponse.Headers, smtpResponse.Body)
	os.Exit(0)
}
