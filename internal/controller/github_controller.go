package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/juliofilizzola/github-discord-bot/internal/model"
	"github.com/juliofilizzola/github-discord-bot/internal/service"
	"net/http"
	"time"
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get repository details"})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"repository": existing})
}

func (c *GitHubController) SaveRepositoryDetails(ctx *gin.Context) {
	var body model.GitHubPullRequestEvent
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if body.PullRequest.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Pull Request ID is required"})
		return
	}

	if body.PullRequest.CreatedAt.IsZero() {
		body.PullRequest.CreatedAt = time.Now()
	}
	if body.PullRequest.UpdatedAt.IsZero() {
		body.PullRequest.UpdatedAt = time.Now()
	}

	if body.PullRequest.User.IdGit == "" && body.PullRequest.User.ID != 0 {
		body.PullRequest.User.IdGit = uuid.New().String()
		body.PullRequest.UserID = body.PullRequest.User.IdGit
	}

	if err := c.service.SaveRepositoryDetails(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save repository details"})
		return
	}
}
