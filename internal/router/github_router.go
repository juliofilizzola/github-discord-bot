package router

import (
	"github-discord-bot/internal/controller"
	"github-discord-bot/internal/service"
	"github.com/gin-gonic/gin"
)

func init() {
	githubService := service.NewGithubService()
	discordService := controller.NewGithubController(githubService)
	RegisterRoutes(func(r *gin.Engine) {
		r.POST("/github/webhook", discordService.GetRepositoryDetails)
	})
}
