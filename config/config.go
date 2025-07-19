package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type (
	// Config -.
	Config struct {
		TelegramBout TelegramBot
	}

	// App -.
	TelegramBot struct {
		Token string `env:"TELEGRAM_BOT_TOKEN,required"`
		Debug bool   `env:"TELEGRAM_BOT_DEBUG,required"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
