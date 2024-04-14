package emailnotificator

import "fmt"

type EmailNotificator struct{}

func (n EmailNotificator) SendAlert() error {
	fmt.Println("Реализация телеграмм нотификатора")
	return nil
}
