package main

import (
	"github.com/juliofilizzola/github-discord-bot/internal/db"
	"github.com/juliofilizzola/github-discord-bot/internal/router"
)

func main() {
	if err := db.InitializeDatabase(); err != nil {
		panic(err)
	}
	router.Init()
}
