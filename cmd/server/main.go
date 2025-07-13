package main

import (
	"github-discord-bot/internal/config"
	"github-discord-bot/internal/db"
)

func main() {
	cfg := config.Load()

	db.Connect(cfg.DatabaseURL)
	//db.DB.AutoMigrate()
}
