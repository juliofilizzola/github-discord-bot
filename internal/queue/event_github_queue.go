package queue

import "github.com/juliofilizzola/github-discord-bot/internal/model"

var EventGithub = make(chan *model.GitHubEvent, 100)

func ConsumeEventGithub() {
	println("ConsumeEventGithub")
	for event := range EventGithub {
		println("Recebido evento do GitHub:")
		if event != nil {
			println("ID:", event)
		} else {
			println("Evento nulo recebido")
		}
	}
}
