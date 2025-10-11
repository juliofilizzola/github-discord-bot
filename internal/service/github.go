package service

import (
	"github.com/juliofilizzola/github-discord-bot/internal/model"
	"github.com/juliofilizzola/github-discord-bot/internal/queue"
	"github.com/juliofilizzola/github-discord-bot/internal/repository"
)

type GitHubService struct {
	repository *repository.GitHubRepository
}

func NewGithubService() *GitHubService {
	return &GitHubService{
		repository: repository.NewGitHubRepository(),
	}
}

func (service *GitHubService) GetRepositoryDetails(owner, repo string) (string, error) {
	existing, err := service.repository.GetRepositoryDetails(owner, repo)
	if err != nil {
		return "", err
	}
	return existing, nil
}

func (service *GitHubService) SaveRepositoryDetails(event *model.GitHubEvent) error {
	if err := service.repository.SaveRepositoryDetails(event); err != nil {
		return err
	}
	queue.EventGithub <- event
	return nil
}
