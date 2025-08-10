package service

import (
	"strconv"
	"time"

	"github.com/juliofilizzola/github-discord-bot/internal/model"

	"github.com/bwmarrin/discordgo"
)

func FormatEmbedDiscord(githubDomain *model.GitHubEvent) discordgo.WebhookParams {
	var reviews []string

	embed := &discordgo.MessageEmbed{
		URL:         githubDomain.PullRequest.HTMLURL,
		Type:        discordgo.EmbedTypeLink,
		Title:       githubDomain.PullRequest.Title,
		Description: "",
		Timestamp:   time.Now().Format(githubDomain.PullRequest.CreatedAt.Format(time.RFC3339)),
		Color:       0,
		Footer: &discordgo.MessageEmbedFooter{
			Text:         "Pull Request #" + strconv.Itoa(githubDomain.PullRequest.Number) + " - " + githubDomain.Repository.Name,
			IconURL:      githubDomain.PullRequest.User.AvatarURL,
			ProxyIconURL: "",
		},
		Image: &discordgo.MessageEmbedImage{
			URL:      githubDomain.Sender.AvatarURL,
			ProxyURL: "",
			Width:    280,
			Height:   20,
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL:      githubDomain.Sender.AvatarURL,
			ProxyURL: "",
			Width:    280,
			Height:   20,
		},
		Author: &discordgo.MessageEmbedAuthor{
			URL:          githubDomain.PullRequest.User.HTMLURL,
			Name:         githubDomain.PullRequest.User.Login,
			IconURL:      githubDomain.PullRequest.User.AvatarURL,
			ProxyIconURL: "",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Branch:",
				Value:  githubDomain.PullRequest.Head.Ref,
				Inline: false,
			},
			{
				Name:   "Merge into:",
				Value:  githubDomain.PullRequest.Base.Ref + " from " + githubDomain.PullRequest.Head.Ref,
				Inline: false,
			},
			{
				Name:   "Status:",
				Value:  githubDomain.PullRequest.State,
				Inline: false,
			},
			{
				Name: "Assinado:",
				Value: func() string {
					if len(githubDomain.PullRequest.Assignee.Login) == 0 {
						return "Não teve assinatura"
					}
					return githubDomain.PullRequest.Assignee.Login
				}(),
				Inline: false,
			},
			{
				Name:   "Codigo adicionado:",
				Value:  strconv.Itoa(githubDomain.PullRequest.Additions),
				Inline: true,
			},
			{
				Name:   "Codigo deletado",
				Value:  strconv.Itoa(githubDomain.PullRequest.Deletions),
				Inline: true,
			},
			{
				Name:   "Commits:",
				Value:  "[commits](" + githubDomain.PullRequest.HTMLURL + "/commits)",
				Inline: false,
			},
			{
				Name:   "Reviews",
				Value:  returnString(reviews),
				Inline: false,
			},
		},
	}

	return discordgo.WebhookParams{
		Content:    "Nova PR no Repositorio: " + githubDomain.Repository.Name,
		Username:   `julio filizzola`,
		AvatarURL:  `julio filizzola`,
		TTS:        false,
		Files:      nil,
		Components: nil,
		Embeds:     []*discordgo.MessageEmbed{embed},
		AllowedMentions: &discordgo.MessageAllowedMentions{
			Parse: []discordgo.AllowedMentionType{
				discordgo.AllowedMentionTypeEveryone,
			},
			Roles:       nil,
			Users:       nil,
			RepliedUser: false,
		},
		Flags: 0,
	}
}

func returnString(reviews []string) string {
	var test string
	if len(reviews) == 0 {
		return "Sem reviews"
	}

	for _, value := range reviews {
		test += value
	}

	return test
}
