package domain

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type TelegramMessage struct {
	ID   int
	Text string
	User TelegramUser
	Chat TelegramChat
}

type TelegramUser struct {
	ID        int64
	FirstName string
	LastName  string
	UserName  string
}

type TelegramChat struct {
	ID int64
}

func ToDomainMessage(update tgbotapi.Update) TelegramMessage {
	return TelegramMessage{
		ID:   update.Message.MessageID,
		Text: update.Message.Text,
		User: TelegramUser{
			ID:        update.Message.From.ID,
			FirstName: update.Message.From.FirstName,
			LastName:  update.Message.From.LastName,
			UserName:  update.Message.From.UserName,
		},
		Chat: TelegramChat{
			ID: update.Message.Chat.ID,
		},
	}
}
