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
	telegramNotificator.SetChatID(config.TelegramConfig.ChatID)
	telegramNotificator.SetToken(config.TelegramConfig.Token)

	discordNotificator := discordNotificator.DiscordNotificator{}
	discordNotificator.SetChatID(config.Discord.ChatID)
	discordNotificator.SetToken(config.Discord.Token)

	tg := telegramNotificator.Create(config.TelegramConfig.Token, config.TelegramConfig.ChatID)

	notGroup := notificator.Create()
	notGroup.Append(telegramNotificator, discordNotificator, tg)
	notGroup.SendAlerts("test group notifications func")
}
