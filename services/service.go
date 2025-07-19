package services

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strings"

	"github.com/SHshzik/genesys_helper/domain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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

func (s *Service) RollDice(update tgbotapi.Update) {
	// Разделяем строку по двоеточию
	parts := strings.Split(update.Message.Text, ":")

	// Получаем вторую часть (ABC)
	diceCommand := parts[1]
	dices := splitByPattern(diceCommand)

	log.Println(dices)

	successCount := 0
	advantageCount := 0
	values := domain.BonusDice[rand.Intn(6)+1]
	for _, value := range values {
		switch value {
		case domain.Success:
			successCount += 1
		case domain.Advantage:
			advantageCount += 1
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
	re := regexp.MustCompile(`\d?[A-Z]`)
	return re.FindAllString(s, -1)
}
