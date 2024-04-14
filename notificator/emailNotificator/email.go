package emailnotificator

import "fmt"

type EmailNotificator struct{}

func (n EmailNotificator) SendAlert(a ...any) error {
	fmt.Println("Реализация email нотификатора")
	return nil
}
