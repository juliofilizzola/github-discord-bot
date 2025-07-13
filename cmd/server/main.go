package main

import (
	"github-discord-bot/internal/db"
	"github-discord-bot/internal/router"
)

func main() {
	if err := db.InitializeDatabase(); err != nil {
		panic(err)
	}
	router.Init()
}
