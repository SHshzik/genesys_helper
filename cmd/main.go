package main

import (
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SHshzik/genesys_helper/adapters/sqlite_adapter"
	"github.com/SHshzik/genesys_helper/config"
	"github.com/SHshzik/genesys_helper/handlers"
	"github.com/SHshzik/genesys_helper/pkg/logger"
	"github.com/SHshzik/genesys_helper/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Panic(err)
	}

	l := logger.New(config.Log.Level)

	db, err := sql.Open("sqlite3", "file:genesys_helper.db?cache=shared&mode=rwc")
	if err != nil {
		log.Panic(err)
	}

	sqlite := sqlite_adapter.NewSqliteAdapter(db, l)

	bot, err := tgbotapi.NewBotAPI(config.TelegramBout.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = config.TelegramBout.Debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	service := services.NewService(bot, sqlite, l)

	botHandler := handlers.NewBot(bot, u, service, l)

	go botHandler.Listen()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	s := <-interrupt
	log.Printf("app - Run - signal: " + s.String())
}
