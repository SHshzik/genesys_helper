package services

import (
	"fmt"
	"strings"

	"github.com/SHshzik/genesys_helper/domain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *Service) CharacterInfo(message domain.TelegramMessage) {
	user, err := s.GetOrCreateUser(message.User)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Ошибка при получении пользователя")
		s.bot.Send(msg)
		return
	}

	character, err := s.sqliteAdapter.GetCharacterByUserID(user.ID)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Ошибка при получении персонажа")
		s.bot.Send(msg)
		return
	}

	response := strings.Builder{}
	response.WriteString("Информация о персонаже:")
	response.WriteString("\n")

	response.WriteString(fmt.Sprintf("Имя: %s", character.Name))
	response.WriteString("\n")

	msg := tgbotapi.NewMessage(message.Chat.ID, response.String())
	s.bot.Send(msg)
}
