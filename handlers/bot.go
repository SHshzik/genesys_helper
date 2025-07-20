package handlers

import (
	"strings"

	"github.com/SHshzik/genesys_helper/domain"
	"github.com/SHshzik/genesys_helper/pkg/logger"
	"github.com/SHshzik/genesys_helper/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	uConfig tgbotapi.UpdateConfig
	service *services.Service
	l       logger.Interface
}

func NewBot(bot *tgbotapi.BotAPI, uConfig tgbotapi.UpdateConfig, service *services.Service, l logger.Interface) *Bot {
	return &Bot{bot: bot, uConfig: uConfig, service: service, l: l}
}

func (b *Bot) Listen() {
	updates := b.bot.GetUpdatesChan(b.uConfig)

	for update := range updates {
		message := domain.ToDomainMessage(update)
		switch {
		case message.Text == "/start":
			b.Start(message)
		case strings.HasPrefix(message.Text, "/roll"):
			b.RollDice(message)
		case message.Text == "/character":
			b.CharacterInfo(message)
		case message.Text == "/info":
			b.Info(message)
		case strings.HasPrefix(message.Text, "/set_name"):
			b.SetName(message)
		}
	}
}

func (b *Bot) Start(message domain.TelegramMessage) {
	b.service.Start(message)
}

func (b *Bot) RollDice(message domain.TelegramMessage) {
	b.service.RollDice(message)
}

func (b *Bot) CharacterInfo(message domain.TelegramMessage) {
	b.service.CharacterInfo(message)
}

func (b *Bot) Info(message domain.TelegramMessage) {
	b.service.Info(message)
}

func (b *Bot) SetName(message domain.TelegramMessage) {
	b.service.SetName(message)
}
