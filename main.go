package main

import (
	"github.com/LittleDrongo/fmn-notificator/notificator"
	emailnotificator "github.com/LittleDrongo/fmn-notificator/notificator/emailNotificator"
	"github.com/LittleDrongo/fmn-notificator/notificator/telegramNotificator"
)

func main() {
	NotificatorsSampleUse()

}

func NotificatorsSampleUse() {

	telegramNotificator := telegramNotificator.TelegramNotificator{}
	emailNotificators := emailnotificator.EmailNotificator{}

	// fmt.Println(teleNotif, emailNotif)

	_ = telegramNotificator.SendAlert()
	_ = emailNotificators.SendAlert()

	arrayNotif := []notificator.Notificator{
		telegramNotificator,
		emailNotificators,
	}

	for _, n := range arrayNotif {

		n.SendAlert()

	}

}
