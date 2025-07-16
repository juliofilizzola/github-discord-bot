package router

import (
	"github.com/gin-gonic/gin"
	"github.com/juliofilizzola/github-discord-bot/internal/config"
)

var routerRegistrations []func(*gin.Engine)

func RegisterRoutes(funcCallback func(r *gin.Engine)) {
	routerRegistrations = append(routerRegistrations, funcCallback)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
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
