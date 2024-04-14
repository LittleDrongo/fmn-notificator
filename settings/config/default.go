package config

import "github.com/LittleDrongo/fmn-notificator/settings"

var TelegramConfigDefault = settings.TelegramSettings{
	ChatID: "",
	Token:  "",
}

var DiscordConfigDefault = settings.DiscordSettings{
	ChatID: "",
	Token:  "",
}
