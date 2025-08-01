package queue

import "github.com/juliofilizzola/github-discord-bot/internal/model"

var EventGithub = make(chan *model.GitHubPullRequestEvent, 100)

func ConsumeEventGithub() {
	for event := range EventGithub {
		println(event)
	}
}
