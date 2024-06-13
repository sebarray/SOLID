package service

import "fmt"

type Notifier interface {
	Send(message string)
}

// EmailSender es una implementación concreta de Notifier que envía correos electrónicos.
type EmailSender struct{}

func (EmailSender) Send(message string) {
	fmt.Println("Sending email with message:", message)
}

// SmsSender es una implementación concreta de Notifier que envía SMS.
type SmsSender struct{}

func (SmsSender) Send(message string) {
	fmt.Println("Sending SMS with message:", message)
}
