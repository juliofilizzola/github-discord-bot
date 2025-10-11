package router

import (
	"github.com/gin-gonic/gin"
	"github.com/juliofilizzola/github-discord-bot/internal/config"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var routerRegistrations []func(*gin.Engine)

func RegisterRoutes(funcCallback func(r *gin.Engine)) {
	routerRegistrations = append(routerRegistrations, funcCallback)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
	for _, register := range routerRegistrations {
		register(router)
	}

	return router
}

func Init() {
	cfg := config.Load()
	router := SetupRouter()

	err := router.Run(cfg.Port)
	if err != nil {
		return
	}
}
