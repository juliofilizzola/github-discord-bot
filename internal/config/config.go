package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL           string
	DatabaseUrlMigrations string
	DiscordWebhookURL     string
	DiscordToken          string
	Port                  string
}

var (
	cfg  *Config
	once sync.Once
)

func Load() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		cfg = &Config{
			DatabaseURL:           os.Getenv("DATABASE_URL"),
			DiscordWebhookURL:     os.Getenv("DISCORD_WEBHOOK_URL"),
			Port:                  os.Getenv("PORT"),
			DatabaseUrlMigrations: os.Getenv("DATABASE_URL_MIGRATIONS"),
			DiscordToken:          os.Getenv("DISCORD_TOKEN"),
		}

		if cfg.DatabaseURL == "" || cfg.DiscordWebhookURL == "" || cfg.Port == "" || cfg.DiscordToken == "" {
			log.Fatal("Missing required environment variables")
		}
	})

	return cfg
}
