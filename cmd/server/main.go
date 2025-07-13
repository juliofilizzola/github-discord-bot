package main

import (
	"fmt"
	"github-discord-bot/internal/config"
	"github-discord-bot/internal/db"
	"github-discord-bot/internal/model"
)

func main() {
	cfg := config.Load()

	db.Connect(cfg.DatabaseURL)
	err := db.DB.AutoMigrate(&model.GitHubEvent{}, &model.User{}, &model.Repository{}, &model.PullRequest{})
	if err != nil {
		fmt.Printf("erro ao migrar banco de dados: %v\n", err)
	}
}
