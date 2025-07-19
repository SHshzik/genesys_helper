package handlers

import (
	"strings"

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
		switch {
		case update.Message.Text == "start" || update.Message.Text == "старт":
			b.Start(update)
		case strings.HasPrefix(update.Message.Text, "G:"):
			b.RollDice(update)
		}
	}
}

func (b *Bot) Start(update tgbotapi.Update) {
	b.service.Start(update)
}

func (b *Bot) RollDice(update tgbotapi.Update) {
	b.service.RollDice(update)
}
