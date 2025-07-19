package services

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Service struct {
	bot *tgbotapi.BotAPI
}

func NewService(bot *tgbotapi.BotAPI) *Service {
	return &Service{bot: bot}
}

func (s *Service) Start(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, я Genesys Helper!")
	// msg.ReplyToMessageID = update.Message.MessageID

	s.bot.Send(msg)
}
