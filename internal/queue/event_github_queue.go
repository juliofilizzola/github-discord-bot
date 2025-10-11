package queue

import (
	"github.com/juliofilizzola/github-discord-bot/internal/discord"
	"github.com/juliofilizzola/github-discord-bot/internal/model"
)

var EventGithub = make(chan *model.GitHubEvent, 100)

func ConsumeEventGithub() {
	println("ConsumeEventGithub")
	for event := range EventGithub {
		println("Recebido evento do GitHub:")
		if event != nil {
			if err := discord.SendEmbedToDiscord("", event); err != nil {
				println("Erro ao enviar embed para Discord:", err.Error())
			}
		} else {
			println("Evento nulo recebido")
		}
	}
}
