package service

import (
	"github.com/juliofilizzola/github-discord-bot/internal/config"

	"github.com/bwmarrin/discordgo"
	"github.com/juliofilizzola/github-discord-bot/internal/model"
)

var (
	DiscordServer *discordgo.Session
)

func SendEmbedToDiscord(webhookId string, event *model.GitHubEvent) error {
	embed := FormatEmbedDiscord(event)
	cfg := config.Load()
	token := cfg.DiscordToken
	_, err := DiscordServer.WebhookExecute(webhookId, token, false, &embed)
	if err != nil {
		println("Error sending embed to Discord:", err.Error())
		return err
	}
	return nil
}
