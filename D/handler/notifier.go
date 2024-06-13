package handler

import "d/service"

type NotificationHandler struct {
	Notifier service.Notifier
}

func (n NotificationHandler) Notify(message string) {
	n.Notifier.Send(message)
}
