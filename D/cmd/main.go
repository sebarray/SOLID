package d

import (
	"d/handler"
	"d/service"
)

//	que carajos es inversion de dependencias?
//Los módulos de alto nivel no deben depender de módulos de bajo nivel. Ambos deben depender de abstracciones.
//Las abstracciones no deben depender de detalles. Los detalles deben depender de abstracciones

func main() {
	// Utilizando el EmailSender
	emailNotifier := service.EmailSender{}
	notificationService := handler.NotificationHandler{Notifier: emailNotifier}
	notificationService.Notify("Hello via Email!")

	// Utilizando el SmsSender
	smsNotifier := service.SmsSender{}
	notificationService = handler.NotificationHandler{Notifier: smsNotifier}
	notificationService.Notify("Hello via SMS!")
}
