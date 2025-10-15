package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/juliofilizzola/github-discord-bot/internal/config"
	"github.com/juliofilizzola/github-discord-bot/internal/model"
	"github.com/juliofilizzola/github-discord-bot/internal/utils"
)

var (
	DiscordServer *discordgo.Session
)

func DiscordConfig() (*discordgo.Session, error) {
	cfg := config.Load()
	discord, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {

		println("Error creating Discord session:")
		return discord, err
	}
	return discord, nil
}

func SendEmbedToDiscord(webhookId string, event *model.WebhookEvent) error {
	embed := utils.FormatEmbedDiscord(event)
	cfg := config.Load()
	token := cfg.DiscordToken

	s, err := DiscordConfig()
	if err != nil {
		return err
	}
	_, err = s.WebhookExecute(webhookId, token, false, &embed)
	if err != nil {
		println("Error sending embed to Discord:", err.Error())
		return err
	}
	return nil
}
