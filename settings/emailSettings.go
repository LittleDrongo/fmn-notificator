package settings

import (
	emailnotificator "github.com/LittleDrongo/fmn-notificator/notificator/emailNotificator"
	"gopkg.in/gomail.v2"
)

type EmailNotificatorSettings struct {
	Dialer   gomail.Dialer
	Settings emailnotificator.EmailSettings
}
