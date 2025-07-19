package service

import (
	"bytes"
	"net/http"
	"os"
)

func SendToDiscord(message string) (err error) {
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")
	payload := []byte(`{"content": "` + message + `"}`)
	_, err = http.Post(webhookURL, "application/json", bytes.NewBuffer(payload))
	return err
}
