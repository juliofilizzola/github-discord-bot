package repository

import (
	"github.com/juliofilizzola/github-discord-bot/internal/db"
	"github.com/juliofilizzola/github-discord-bot/internal/model"
	"gorm.io/gorm"
)

type GitHubRepository struct{}

func NewGitHubRepository() *GitHubRepository {
	return &GitHubRepository{}
}

func (r *GitHubRepository) GetUserOwner(ownerId string) (*model.GitHubUser, error) {
	var existing model.GitHubUser
	if err := db.GetDB().Where("id = ?", ownerId).First(&existing).Error; err != nil {
		return nil, err
	}
	return &existing, nil
}

func (r *GitHubRepository) GetRepositoryDetails(owner, repo string) (string, error) {
	var existing model.GitHubEvent
	if err := db.GetDB().Where("owner = ? AND repo = ?", owner, repo).First(&existing).Error; err != nil {
		return "", err
	}
	return existing.ID, nil
}

func (r *GitHubRepository) SaveRepositoryDetails(event *model.GitHubEvent) error {
	if err := db.GetDB().Session(&gorm.Session{FullSaveAssociations: true}).Create(&event).Error; err != nil {
		println(err.Error())
		return err
	}
	return nil
}
