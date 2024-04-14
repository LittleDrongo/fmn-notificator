package notificator

type Notificator interface {
	SendAlert(a ...any) error
}

func SendAlertNotification(n Notificator) error {
	return n.SendAlert()
}
