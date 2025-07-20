package services

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/SHshzik/genesys_helper/domain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *Service) RollDice(message domain.TelegramMessage) {
	var successCount int
	var advantageCount int
	var triumphCount int
	var failureCount int
	var complicationCount int
	var crashCount int

	parts := strings.Split(message.Text, diceRollCommandDelimiter)
	diceCommand := parts[1]
	dices := splitByPattern(diceCommand)
	tokens, err := parseTokens(dices)
	if err != nil {
		msg := tgbotapi.NewMessage(message.User.ID, "Неизвестный кубик")
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

	msg := tgbotapi.NewMessage(message.Chat.ID, response.String())
	s.bot.Send(msg)
}
