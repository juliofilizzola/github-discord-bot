package repository

import (
	"github.com/google/uuid"
	"github.com/juliofilizzola/github-discord-bot/internal/db"
	"github.com/juliofilizzola/github-discord-bot/internal/model"
	"gorm.io/gorm"
)

type GitHubRepository struct{}

func NewGitHubRepository() *GitHubRepository {
	return &GitHubRepository{}
}

func (r *GitHubRepository) GetRepositoryDetails(owner, repo string) (string, error) {
	var existing model.GitHubPullRequestEvent
	if err := db.DB.Where("owner = ? AND repo = ?", owner, repo).First(&existing).Error; err != nil {
		return "", err // Outro erro
	}
	return existing.ID, nil
}

func (r *GitHubRepository) SaveRepositoryDetails(event *model.GitHubPullRequestEvent) error {
	if event.PullRequest.IdPullRequest == "" {
		event.PullRequest.IdPullRequest = uuid.New().String()
		event.PullRequestID = event.PullRequest.IdPullRequest
	}

	if err := db.DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(&event).Error; err != nil {
		return err
	}
	return nil
}
