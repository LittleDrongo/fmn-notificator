package notificator

import (
	"sync"

	"github.com/LittleDrongo/fmn-lib/exception"
)

type NotificatorsGroup struct {
	List []Notificator
}

func Create() NotificatorsGroup {
	return NotificatorsGroup{}
}

func (ng *NotificatorsGroup) Append(notificators ...Notificator) {
	ng.List = append(ng.List, notificators...)
}

func (ng *NotificatorsGroup) SendAlerts(a ...any) {

	notificators := ng.List
	var wg sync.WaitGroup
	wg.Add(len(notificators))

	for _, n := range notificators {

		go func(n Notificator) {
			defer wg.Done()
			err := n.SendAlert(a...)
			exception.Println(err, "ошибка: SendAlerts групповой отправки нотификатора:")
		}(n)
	}

	wg.Wait()

}
