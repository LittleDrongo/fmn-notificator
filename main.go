package main

import (
	"sync"

	"github.com/LittleDrongo/fmn-lib/exception"
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

	notificators := []notificator.Notificator{
		telegramNotificator,
		discordNotificator,
	}

	var wg sync.WaitGroup
	wg.Add(len(notificators))

	for _, n := range notificators {
		// Запускаем каждый вызов SendAlert в отдельной горутине
		go func(n notificator.Notificator) {
			defer wg.Done() // Уменьшаем счетчик горутин при завершении
			err := n.SendAlert("Проверка асинхронного выполнения кода 2")
			exception.Println(err, "Ошибка нотификатора:")
		}(n)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
}

// func NotificatorsSampleUse() {

// 	telegramNotificator := telegramNotificator.TelegramNotificator{}
// 	telegramNotificator.SetChatID(config.TelegramConfig.ChatID)
// 	telegramNotificator.SetToken(config.TelegramConfig.Token)

// 	discordNotificator := discordNotificator.DiscordNotificator{}
// 	discordNotificator.SetChatID(config.Discord.ChatID)
// 	discordNotificator.SetToken(config.Discord.Token)

// 	notificators := []notificator.Notificator{
// 		telegramNotificator,
// 		discordNotificator,
// 	}

// 	for _, n := range notificators {
// 		// Запускаем каждый вызов SendAlert в отдельной горутине
// 		go func(n notificator.Notificator) {
// 			err := n.SendAlert("Проверка асинхронного выполнения кода")
// 			exception.Println(err, "Ошибка нотификатора:")
// 		}(n)
// 	}

// }
