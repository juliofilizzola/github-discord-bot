package model

type GitHubEvent struct {
	ID            string `gorm:"primaryKey"`
	IdPullRequest string `json:"IdPullRequest"`
	PullRequest   string `gorm:"foreignKey:IdPullRequest;references:ID"`
	status        string
}

func init() {
	RegisterModel(&GitHubEvent{})
}
