package services

import (
	"fmt"
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
	var triumphCount int
	var failureCount int
	var complicationCount int
	var crashCount int

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
		var count int
		switch token.Letter {
		case domain.BonusDiceLetter:
			count = 6
			dice = domain.BonusDice
		case domain.AbilityDiceLetter:
			count = 8
			dice = domain.AbilityDice
		case domain.MasterDiceLetter:
			count = 12
			dice = domain.MasterDice
		case domain.PenaltyDiceLetter:
			count = 6
			dice = domain.PenaltyDice
		case domain.DifficultyDiceLetter:
			count = 8
			dice = domain.DifficultyDice
		case domain.ChallengeDiceLetter:
			count = 12
			dice = domain.ChallengeDice
		}

		for i := 0; i < token.Count; i++ {
			random := rand.Intn(count) + 1
			values := dice[random]
			for _, value := range values {
				switch value {
				case domain.Success:
					successCount += 1
				case domain.Advantage:
					advantageCount += 1
				case domain.Triumph:
					successCount += 1
					triumphCount += 1
				case domain.Failure:
					failureCount += 1
				case domain.Complication:
					complicationCount += 1
				case domain.Crash:
					failureCount += 1
					crashCount += 1
				}
			}
		}
	}

	response := strings.Builder{}
	response.WriteString("Результат броска: \n")

	var resultString string
	result := successCount - failureCount
	if result <= 0 {
		resultString = "Неудача"
	} else {
		resultString = "Успех"
	}
	response.WriteString(fmt.Sprintf("Успех: %s", resultString))
	response.WriteString("\n")

	advantageScore := advantageCount - complicationCount
	if advantageScore < 0 {
		response.WriteString(fmt.Sprintf("Осложнение: %d", -advantageScore))
		response.WriteString("\n")
	} else if advantageScore > 0 {
		response.WriteString(fmt.Sprintf("Преимущество: %d", advantageScore))
		response.WriteString("\n")
	}

	if triumphCount > 0 {
		response.WriteString(fmt.Sprintf("Триумф: %d", triumphCount))
		response.WriteString("\n")
	}

	if crashCount > 0 {
		response.WriteString(fmt.Sprintf("Крах: %d", crashCount))
		response.WriteString("\n")
	}

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
