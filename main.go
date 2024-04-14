package main

import (
	"github.com/LittleDrongo/fmn-notificator/notificator"
	"github.com/LittleDrongo/fmn-notificator/notificator/discordNotificator"
	"github.com/LittleDrongo/fmn-notificator/notificator/telegramNotificator"
	"github.com/LittleDrongo/fmn-notificator/settings/config"
)

func main() {
	NotificatorsSampleUse()

}

func NotificatorsSampleUse() {
	telegramNotificator := telegramNotificator.TelegramNotificator{}
	telegramNotificator.SetChatID(config.Telegram.ChatID)
	telegramNotificator.SetToken(config.Telegram.Token)

	discordNotificator := discordNotificator.DiscordNotificator{}
	discordNotificator.SetChatID(config.Discord.ChatID)
	discordNotificator.SetToken(config.Discord.Token)

	tg := telegramNotificator.Create(config.Telegram.Token, config.Telegram.ChatID)

	notGroup := notificator.Create()
	notGroup.Append(telegramNotificator, discordNotificator, tg)
	notGroup.SendAlerts("hello text")
}
