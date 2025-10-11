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
		api := r.Group("/api/v1/github")
		api.GET("/github/webhook", discordService.GetRepositoryDetails)
		api.POST("/github", discordService.SaveRepositoryDetails)
	})
}
