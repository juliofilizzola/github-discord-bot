package model

import "time"

type PullRequest struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	GitHubID     int64      `gorm:"uniqueIndex" json:"github_id"`
	Number       int        `json:"number"`
	Title        string     `json:"title"`
	State        string     `json:"state"`
	HtmlURL      string     `json:"html_url"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	MergedAt     *time.Time `json:"merged_at"`
	ClosedAt     *time.Time `json:"closed_at"`
	HeadRef      string     `json:"head_ref"`
	HeadSha      string     `json:"head_sha"`
	BaseRef      string     `json:"base_ref"`
	Commits      int        `json:"commits"`
	Additions    int        `json:"additions"`
	Deletions    int        `json:"deletions"`
	ChangedFiles int        `json:"changed_files"`

	RepositoryID uint       `json:"repository_id"`
	Repository   Repository `gorm:"foreignKey:RepositoryID" json:"repository"`

	AuthorID uint `json:"author_id"`
	Author   User `gorm:"foreignKey:AuthorID" json:"author"`
}
