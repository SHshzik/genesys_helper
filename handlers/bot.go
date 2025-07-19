package handlers

import (
	"strings"

	"github.com/SHshzik/genesys_helper/domain"
	"github.com/SHshzik/genesys_helper/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	uConfig tgbotapi.UpdateConfig
	service *services.Service
}

func NewBot(bot *tgbotapi.BotAPI, uConfig tgbotapi.UpdateConfig, service *services.Service) *Bot {
	return &Bot{bot: bot, uConfig: uConfig, service: service}
}

func (b *Bot) Listen() {
	updates := b.bot.GetUpdatesChan(b.uConfig)

	for update := range updates {
		message := domain.ToDomainMessage(update)
		switch {
		case message.Text == "start" || message.Text == "старт":
			b.Start(message)
		case strings.HasPrefix(message.Text, "G:"):
			b.RollDice(message)
		case message.Text == "лист персонажей":
			b.ListCharacters(message)
		}
	}
}

func (b *Bot) Start(message domain.TelegramMessage) {
	b.service.Start(message)
}

func (b *Bot) RollDice(message domain.TelegramMessage) {
	b.service.RollDice(message)
}

func (b *Bot) ListCharacters(message domain.TelegramMessage) {
	b.service.ListCharacters(message)
}
