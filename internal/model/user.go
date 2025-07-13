package model

import "time"

type User struct {
	ID        string `gorm:"primaryKey" json:"id"`
	GitHubID  int64  `gorm:"uniqueIndex" json:"github_id"`
	Login     string `json:"login"`
	HtmlURL   string `json:"html_url"`
	AvatarURL string `json:"avatar_url"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func init() {
	RegisterModel(&User{})
}
