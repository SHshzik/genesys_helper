package services

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/SHshzik/genesys_helper/domain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	diceRollCommandDelimiter = ":"
)

type Service struct {
	bot *tgbotapi.BotAPI
}

func NewService(bot *tgbotapi.BotAPI) *Service {
	return &Service{bot: bot}
}

func (s *Service) Start(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, я Genesys Helper!")
	msg.ReplyToMessageID = update.Message.MessageID

	s.bot.Send(msg)
}

func (s *Service) RollDice(update tgbotapi.Update) {
	var successCount int
	var advantageCount int

	parts := strings.Split(update.Message.Text, diceRollCommandDelimiter)
	diceCommand := parts[1]
	dices := splitByPattern(diceCommand)
	tokens, err := parseTokens(dices)
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестный кубик")
		s.bot.Send(msg)

		return
	}

	for _, token := range tokens {
		var dice domain.Dice
		switch token.Letter {
		case domain.BonusDiceLetter:
			dice = domain.BonusDice
		}

		for i := 0; i < token.Count; i++ {
			random := rand.Intn(6) + 1
			log.Println(random)
			values := dice[random]
			log.Println(values)
			for _, value := range values {
				switch value {
				case domain.Success:
					successCount += 1
				case domain.Advantage:
					advantageCount += 1
				}
			}
		}
	}

	response := strings.Builder{}
	response.WriteString("Бросок кости бонуса: \n")
	response.WriteString(fmt.Sprintf("Успех: %d", successCount))
	response.WriteString("\n")
	response.WriteString(fmt.Sprintf("Преимущество: %d", advantageCount))
	response.WriteString("\n")

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, response.String())
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
				fmt.Printf("ошибка разбора числа: %v\n", err)
				continue
			}
		}

		tokens = append(tokens, domain.Token{Count: count, Letter: letter})
	}
	return tokens, nil
}
