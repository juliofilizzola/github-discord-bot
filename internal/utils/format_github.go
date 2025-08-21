package utils

import (
	"github.com/bwmarrin/discordgo"
	"github.com/juliofilizzola/github-discord-bot/internal/model"
)

func FormatGithub(event *model.GitHubEvent) discordgo.WebhookParams {
	if event == nil {
		return discordgo.WebhookParams{}
	}

	pr := event.PullRequest
	repo := event.Repository
	sender := event.Sender
	action := event.Action

	fields := []*discordgo.MessageEmbedField{
		createField("Branch:", pr.Head.Ref, false),
		createField("Merge into:", pr.Base.Ref, false),
		createField("Status:", pr.State, false),
		createField("Author:", pr.User.Login, false),
		createField("Repository:", repo.Name, false),
		createField("Action:", action, false),
	}

	if action == "opened" {
		fields = append(fields, createField("Assignee:", pr.Assignee.Login, false))
	}

}

func createField(name, value string, inline bool) *discordgo.MessageEmbedField {
	return &discordgo.MessageEmbedField{
		Name:   name,
		Value:  value,
		Inline: inline,
	}
}
