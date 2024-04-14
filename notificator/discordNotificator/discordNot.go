package discordNotificator

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

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

// func (n *DiscordNotificator) sendMessage(msg string) error {

// 	// Токен вашего бота Discord
// 	token := n.token

// 	// Создаем новую сессию Discord бота
// 	dg, err := discordgo.New("Bot " + token)
// 	if err != nil {
// 		// fmt.Println("Error creating Discord session: ", err)
// 		return err
// 	}

// 	// Обработчик события готовности
// 	dg.AddHandler(func(s *discordgo.Session, event *discordgo.Ready) {
// 		onReady(s, event, n, msg)
// 	})

// 	// Открываем соединение с Discord
// 	err = dg.Open()
// 	if err != nil {
// 		// fmt.Println("Error opening Discord connection: ", err)
// 		return err
// 	}

// 	dg.Close()
// 	return nil
// }

// func onReady(s *discordgo.Session, event *discordgo.Ready, n *DiscordNotificator, message string) {
// 	// ID канала, в который вы хотите отправить сообщение
// 	channelID := n.chatID

// 	// Текст сообщения
// 	// message := "Hello, Discord!"

// 	// Отправляем сообщение в указанный канал
// 	_, err := s.ChannelMessageSend(channelID, message)
// 	if err != nil {
// 		fmt.Println("Error sending message: ", err)
// 		return
// 	}
// }
