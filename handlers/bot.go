package handlers

import (
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
		switch update.Message.Text {
		case "start", "старт":
			b.Start(update)
		}
	}
}

func (b *Bot) Start(update tgbotapi.Update) {
	b.service.Start(update)
}
