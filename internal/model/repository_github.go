package model

import "time"

type Repository struct {
	ID            string `gorm:"primaryKey" json:"id"`
	GitHubID      int64  `gorm:"uniqueIndex" json:"github_id"`
	Name          string `json:"name"`
	FullName      string `json:"full_name"`
	HtmlURL       string `json:"html_url"`
	Description   string `json:"description"`
	Language      string `json:"language"`
	Private       bool   `json:"private"`
	DefaultBranch string `json:"default_branch"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func init() {
	RegisterModel(&Repository{})
}
