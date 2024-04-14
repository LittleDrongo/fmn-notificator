package discordNotificator

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (d DiscordNotificator) Create(token, chatID string) DiscordNotificator {
	return DiscordNotificator{token: token, chatID: chatID}
}

type DiscordNotificator struct {
	token   string
	chatID  string
	session *discordgo.Session
}

func (n DiscordNotificator) SendAlert(a ...any) error {
	msg := fmt.Sprint(a...)
	err := n.sendMessage(msg)

	if err != nil {
		return err
	}
	return nil
}

func (n *DiscordNotificator) SetToken(token string) {
	n.token = token
}

func (n *DiscordNotificator) SetChatID(chatID string) {
	n.chatID = chatID
}

func (n *DiscordNotificator) sendMessage(msg string) error {
	// Проверяем, есть ли уже у нас сессия Discord
	if n.session == nil {
		var err error
		n.session, err = discordgo.New("Bot " + n.token)
		if err != nil {
			return err
		}
	}

	// ID канала, в который вы хотите отправить сообщение
	channelID := n.chatID

	// Отправляем сообщение в указанный канал
	_, err := n.session.ChannelMessageSend(channelID, msg)
	if err != nil {
		return err
	}

	return nil
}
