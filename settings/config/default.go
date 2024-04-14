package config

import (
	emailnotificator "github.com/LittleDrongo/fmn-notificator/notificator/emailNotificator"
	"github.com/LittleDrongo/fmn-notificator/settings"
	"gopkg.in/gomail.v2"
)

var TelegramConfigDefault = settings.TelegramSettings{
	ChatID: "",
	Token:  "",
}

var DiscordConfigDefault = settings.DiscordSettings{
	ChatID: "",
	Token:  "",
}

var EmailConfigDefault = emailnotificator.EmailSettings{
	From:   "",
	To:     []string{},
	Cc:     []string{},
	Subj:   "",
	Dialer: *gomail.NewDialer("", 000, "", ""),
}
