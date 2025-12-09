package discord

import (
	"fmt"

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

func SendEmbedToDiscord(webhookId string, event *model.GitHubEvent) error {
	embed := utils.FormatEmbedDiscord(event)
	//cfg := config.Load()
	//token := cfg.DiscordToken
	println("Enviando embed para o Discord")
	fmt.Printf("Webhook ID:%+v\n", embed)
	println(webhookId)
	s, err := DiscordConfig()
	if err != nil {
		return err
	}
	_, err = s.WebhookExecute("1119049844602974339", "xEn_Ok9gsHt0kH0LCMFG-oDM_0NdIixRtEdvPx1Yb1_pUTsgxv4kFPt_rc_OAPMNXcDa", false, &embed)
	if err != nil {
		println("Error sending embed to Discord:", err.Error())
		return err
	}
	return nil
}
