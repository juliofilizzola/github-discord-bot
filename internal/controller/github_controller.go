package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/juliofilizzola/github-discord-bot/internal/service"
)

type GitHubController struct {
	service *service.GitHubService
}

func NewGithubController(s *service.GitHubService) *GitHubController {
	return &GitHubController{
		service: s,
	}
}

func (c *GitHubController) GetRepositoryDetails(ctx *gin.Context) {
	owner := ctx.Param("owner")
	repo := ctx.Param("repo")

	existing, err := c.service.GetRepositoryDetails(owner, repo)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to get repository details"})
		return
	}

	ctx.JSON(200, gin.H{"repository": existing})
}
