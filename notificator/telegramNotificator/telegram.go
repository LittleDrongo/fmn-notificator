package telegramNotificator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LittleDrongo/fmn-lib/console/color"
)

func (t TelegramNotificator) Create(token, chatID string) TelegramNotificator {
	return TelegramNotificator{token: token, chatID: chatID}
}

type TelegramNotificator struct {
	token  string
	chatID string
}

func (n TelegramNotificator) SendAlert(a ...any) error {
	msg := fmt.Sprint(a...)
	err := n.sendMessage(msg)

	if err != nil {
		return err
	}
	return nil

}

func (n *TelegramNotificator) SetToken(token string) {
	n.token = token
}

func (n *TelegramNotificator) SetChatID(chatID string) {
	n.chatID = chatID
}

type message struct {
	Text   string `json:"text"`
	ChatID string `json:"chat_id"`
}

func (n TelegramNotificator) sendMessage(msg string) error {
	msg2 := message{
		Text:   msg,
		ChatID: n.chatID,
	}

	fmt.Println(color.RED, msg2, color.RESET)

	msgJSON, err := json.Marshal(msg2)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", n.token)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(msgJSON))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ошибка отправки сообщения в телеграм: статус код: %d", resp.StatusCode)
	}

	return nil
}
