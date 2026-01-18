package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlerGitHubWebhook(c *gin.Context) {
	// todo: Handle GitHub webhook events here
	c.JSON(http.StatusOK, gin.H{"message": "GitHub webhook received"})
}
