package utils

import (
	"strconv"
	"strings"
	"time"

	"github.com/juliofilizzola/github-discord-bot/internal/constants"
	"github.com/juliofilizzola/github-discord-bot/internal/model"

	"github.com/bwmarrin/discordgo"
)

func alertDiscordColor(title string) int {
	titleLower := strings.ToLower(title)

	switch {
	case strings.HasPrefix(titleLower, "feat"):
		return int(constants.ColorSuccess)
	case strings.HasPrefix(titleLower, "fix"):
		return int(constants.ColorWarning)
	case strings.HasPrefix(titleLower, "hotfix"):
		return int(constants.ColorError)
	case strings.HasPrefix(titleLower, "docs"):
		return int(constants.ColorInfo)
	default:
		return int(constants.ColorDefault)
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

func FormatEmbedDiscord(githubDomain *model.GitHubEvent) discordgo.WebhookParams {
	var reviews []string

	embed := &discordgo.MessageEmbed{
		URL:         githubDomain.PullRequest.HTMLURL,
		Type:        discordgo.EmbedTypeLink,
		Title:       githubDomain.PullRequest.Title,
		Description: "",
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       alertDiscordColor(githubDomain.PullRequest.Title),
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
		Video:    nil,
		Provider: nil,
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
						return "Não teve assinatura."
					}
					return githubDomain.PullRequest.Assignee.Login
				}(),
				Inline: false,
			},
			{
				Name:   "Código adicionado:",
				Value:  strconv.Itoa(githubDomain.PullRequest.Additions),
				Inline: true,
			},
			{
				Name:   "Código deletado",
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
		Content:    "Nova PR no Repository: " + githubDomain.Repository.Name,
		Username:   `julio filizzola`,
		AvatarURL:  githubDomain.Sender.AvatarURL,
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
