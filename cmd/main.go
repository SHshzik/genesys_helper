package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SHshzik/genesys_helper/config"
	"github.com/SHshzik/genesys_helper/handlers"
	"github.com/SHshzik/genesys_helper/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(config.TelegramBout.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = config.TelegramBout.Debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	service := services.NewService(bot)

	botHandler := handlers.NewBot(bot, u, service)

	go botHandler.Listen()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	s := <-interrupt
	log.Printf("app - Run - signal: " + s.String())
}
