package services

import (
	"fmt"
	"strings"

	"github.com/SHshzik/genesys_helper/domain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *Service) SetName(message domain.TelegramMessage) {
	character, err := s.GetOrCreateCharacter(message.User)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Ошибка при получении пользователя")
		s.bot.Send(msg)
		return
	}

	name := strings.TrimPrefix(message.Text, "/set_name")
	name = strings.TrimSpace(name)

	character.Name = name
	err = s.sqliteAdapter.UpdateCharacter(&character)
	if err != nil {
		s.l.Error("SetName", "error", err)
		msg := tgbotapi.NewMessage(message.Chat.ID, "Ошибка при обновлении персонажа")
		s.bot.Send(msg)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Имя персонажа обновлено")
	s.bot.Send(msg)
}

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
