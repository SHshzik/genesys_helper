package services

import (
	"github.com/SHshzik/genesys_helper/domain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *Service) Info(message domain.TelegramMessage) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Информация о боте")
	s.bot.Send(msg)
}
