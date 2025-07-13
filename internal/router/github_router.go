package router

import "github.com/gin-gonic/gin"

func init() {
	RegisterRoutes(func(r *gin.Engine) {
		r.POST("/github/webhook", func(c *gin.Context) {})
	})
}
