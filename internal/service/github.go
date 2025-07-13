package service

import "github-discord-bot/internal/repository"

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
		return "", err // Outro erro
	}
	return existing, nil
}
