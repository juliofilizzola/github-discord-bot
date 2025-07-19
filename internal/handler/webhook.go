package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handlerGitHubWebhook(c *gin.Context) {
	// Handle GitHub webhook events here
	c.JSON(http.StatusOK, gin.H{"message": "GitHub webhook received"})
}
