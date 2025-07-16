package repository

import (
	"github.com/juliofilizzola/github-discord-bot/internal/db"
	"github.com/juliofilizzola/github-discord-bot/internal/model"
)

type GitHubRepository struct{}

func NewGitHubRepository() *GitHubRepository {
	return &GitHubRepository{}
}

func (r *GitHubRepository) GetRepositoryDetails(owner, repo string) (string, error) {
	var existing model.GitHubEvent
	if err := db.DB.Where("owner = ? AND repo = ?", owner, repo).First(&existing).Error; err != nil {
		return "", err // Outro erro
	}
	return existing.ID, nil
}
