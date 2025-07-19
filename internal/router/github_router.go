package router

import (
	"github.com/gin-gonic/gin"
	"github.com/juliofilizzola/github-discord-bot/internal/controller"
	"github.com/juliofilizzola/github-discord-bot/internal/service"
)

func init() {
	githubService := service.NewGithubService()
	discordService := controller.NewGithubController(githubService)
	RegisterRoutes(func(r *gin.Engine) {
		r.GET("/github/webhook", discordService.GetRepositoryDetails)
		r.POST("/github", discordService.SaveRepositoryDetails)
	})
}
