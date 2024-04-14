package telegramNotificator

import "fmt"

type TelegramNotificator struct{}

func (n TelegramNotificator) SendAlert() error {
	fmt.Println("Реализация телеграмм нотификатора")
	return nil
}
