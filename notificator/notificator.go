package notificator

type Notificator interface {
	SendAlert() error
}

func SendAlertNotification(n Notificator) error {
	return n.SendAlert()
}
