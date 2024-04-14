package main

import (
	"os"

	"github.com/LittleDrongo/fmn-notificator/notificator"
	"github.com/LittleDrongo/fmn-notificator/notificator/discordNotificator"
	emailnotificator "github.com/LittleDrongo/fmn-notificator/notificator/emailNotificator"
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

	emailNot := emailnotificator.EmailNotificator{}
	emailNot.SetSettings(config.Email.From, config.Email.To, config.Email.Cc, config.Email.Subj)
	emailNot.SetDialer(config.Email.Dialer.Host, config.Email.Dialer.Port, config.Email.Dialer.Username, config.Email.Dialer.Password)

	tg := telegramNotificator.Create(config.Telegram.Token, config.Telegram.ChatID)

	notGroup := notificator.Create()
	notGroup.Append(telegramNotificator, discordNotificator, tg, emailNot)

	_, err := os.Open("data/file.txt")
	if err != nil {
		notGroup.SendAlerts("Внимание! Приложение упало! ошибка:", err)
	}

}
