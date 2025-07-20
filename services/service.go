package services

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"

	"github.com/SHshzik/genesys_helper/adapters/sqlite_adapter"
	"github.com/SHshzik/genesys_helper/domain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	diceRollCommandDelimiter = " "
)

type Service struct {
	bot           *tgbotapi.BotAPI
	sqliteAdapter *sqlite_adapter.SqliteAdapter
}

func NewService(bot *tgbotapi.BotAPI, sqliteAdapter *sqlite_adapter.SqliteAdapter) *Service {
	return &Service{bot: bot, sqliteAdapter: sqliteAdapter}
}

func (s *Service) Start(message domain.TelegramMessage) {
	msg := tgbotapi.NewMessage(message.User.ID, "Привет, я Genesys Helper!")
	msg.ReplyToMessageID = message.ID

	s.bot.Send(msg)
}

func splitByPattern(s string) []string {
	re := regexp.MustCompile(`\d*[A-Z]`)
	return re.FindAllString(s, -1)
}

func parseTokens(parts []string) ([]domain.Token, error) {
	tokens := make([]domain.Token, 0, len(parts))
	for _, part := range parts {
		// Поиск позиции первой буквы
		i := 0
		for ; i < len(part); i++ {
			if part[i] >= 'A' && part[i] <= 'Z' {
				break
			}
		}

		countStr := part[:i]
		letter := part[i:]
		if !slices.Contains(domain.AvailableLetters, letter) {
			return nil, fmt.Errorf("неизвестная буква: %s", letter)
		}

		// Если нет цифры — по умолчанию 1
		count := 1
		if countStr != "" {
			var err error
			count, err = strconv.Atoi(countStr)
			if err != nil {
				continue
			}
		}

		tokens = append(tokens, domain.Token{Count: count, Letter: letter})
	}
	return tokens, nil
}

func (s *Service) GetOrCreateUser(telegramUser domain.TelegramUser) (domain.User, error) {
	user, err := s.sqliteAdapter.GetUserByID(telegramUser.ID)
	if err != nil {
		user = domain.User{
			ID:        telegramUser.ID,
			FirstName: telegramUser.FirstName,
			LastName:  telegramUser.LastName,
			UserName:  telegramUser.UserName,
		}
		err = s.sqliteAdapter.CreateUser(&user)
		if err != nil {
			return domain.User{}, err
		}
	}

	return user, nil
}
