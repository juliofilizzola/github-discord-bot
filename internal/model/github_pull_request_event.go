package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GitHubPullRequestEvent struct {
	ID          string      `gorm:"type:uuid;primaryKey"`
	Action      string      `gorm:"size:20;index"`
	Number      int         `gorm:"index"`
	PullRequest PullRequest `gorm:"embedded;embeddedPrefix:pr_"`
	Repository  Repository  `gorm:"embedded;embeddedPrefix:repo_"`
	Sender      GitHubUser  `gorm:"embedded;embeddedPrefix:sender_"`
	CreatedAt   int64
	UpdatedAt   int64
}

func (e *GitHubPullRequestEvent) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New().String()
	return
}

type PullRequest struct {
	ID        string     `gorm:"column:id;type:uuid"`
	Title     string     `gorm:"size:255"`
	Body      string     `gorm:"type:text"`
	HtmlURL   string     `gorm:"size:255"`
	State     string     `gorm:"size:50"`
	CreatedAt string     `gorm:"size:50"`
	UpdatedAt string     `gorm:"size:50"`
	User      GitHubUser `gorm:"embedded;embeddedPrefix:user_"`
	HeadRef   string     `gorm:"size:100"`
	HeadSHA   string     `gorm:"size:100"`
	BaseRef   string     `gorm:"size:100"`
	BaseSHA   string     `gorm:"size:100"`
}

type GitHubUser struct {
	ID      string `gorm:"column:id;type:uuid"`
	Login   string `gorm:"size:100"`
	HTMLURL string `gorm:"size:255"`
}

type Repository struct {
	ID       string `gorm:"column:id;type:uuid"`
	Name     string `gorm:"size:100"`
	FullName string `gorm:"size:200"`
	HTMLURL  string `gorm:"size:255"`
	Private  bool
	Owner    GitHubUser `gorm:"embedded;embeddedPrefix:owner_"`
}

func Init() {
	RegisterModel(&GitHubPullRequestEvent{})
	RegisterModel(&PullRequest{})
	RegisterModel(&GitHubUser{})
	RegisterModel(&Repository{})
}
