package queue

import (
	"github.com/juliofilizzola/github-discord-bot/internal/discord"
	"github.com/juliofilizzola/github-discord-bot/internal/model"
)

var EventGithub = make(chan *model.WebhookEvent, 100)

func ConsumeEventGithub() {
	println("ConsumeEventGithub")
	for event := range EventGithub {
		println("Recebido evento do GitHub:")
		if event != nil {
			if err := discord.SendEmbedToDiscord("xEn_Ok9gsHt0kH0LCMFG-oDM_0NdIixRtEdvPx1Yb1_pUTsgxv4kFPt_rc_OAPMNXcDa", event); err != nil {
				println("Erro ao enviar embed para Discord:", err.Error())
			}
		} else {
			println("Evento nulo recebido")
		}
	}
}
