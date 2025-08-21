package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/juliofilizzola/github-discord-bot/internal/config"
)

func DiscordConfig() (*discordgo.Session, error) {
	cfg := config.Load()
	discord, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		return discord, err
	}
	return discord, nil
}
