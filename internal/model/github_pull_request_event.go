package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type GitHubUser struct {
	IdGit     string `gorm:"column:id_git;type:uuid;primaryKey"`
	ID        int    `gorm:"column:id;index"`
	NodeID    string `gorm:"column:node_id;size:50"`
	Login     string `gorm:"size:100"`
	AvatarURL string `gorm:"column:avatar_url;size:255"`
	HTMLURL   string `gorm:"column:html_url;size:255"`
	Type      string `gorm:"size:50"`
	SiteAdmin bool   `gorm:"column:site_admin"`
}

type Repository struct {
	IdRepository   string `gorm:"column:id_repository;type:uuid;primaryKey"`
	ID             int    `gorm:"column:id;index"`
	NodeID         string `gorm:"column:node_id;size:50"`
	Name           string `gorm:"size:100"`
	FullName       string `gorm:"column:full_name;size:200"`
	Private        bool
	HTMLURL        string `gorm:"column:html_url;size:255"`
	Description    string `gorm:"type:text"`
	Fork           bool
	Language       string `gorm:"size:50"`
	HasIssues      bool   `gorm:"column:has_issues"`
	HasProjects    bool   `gorm:"column:has_projects"`
	HasWiki        bool   `gorm:"column:has_wiki"`
	HasPages       bool   `gorm:"column:has_pages"`
	HasDiscussions bool   `gorm:"column:has_discussions"`
	Archived       bool
	Disabled       bool
	OwnerID        string     `gorm:"type:uuid"`
	Owner          GitHubUser `gorm:"foreignKey:OwnerID;references:IdGit"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	PushedAt       time.Time `gorm:"column:pushed_at"`
}

type PullRequest struct {
	IdPullRequest  string `gorm:"column:id_pull_request;type:uuid;primaryKey"`
	ID             int    `gorm:"column:id;index"`
	NodeID         string `gorm:"column:node_id;size:50"`
	Number         int
	State          string `gorm:"size:50"`
	Locked         bool
	Title          string  `gorm:"size:255"`
	Body           *string `gorm:"type:text"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ClosedAt       *time.Time
	MergedAt       *time.Time
	MergeCommitSHA *string `gorm:"column:merge_commit_sha;size:100"`
	Draft          bool
	UserID         string     `gorm:"type:uuid"`
	User           GitHubUser `gorm:"foreignKey:UserID;references:IdGit"`
	HTMLURL        string     `gorm:"column:html_url;size:255"`
	HeadRef        string     `gorm:"size:100"`
	HeadSHA        string     `gorm:"size:100"`
	BaseRef        string     `gorm:"size:100"`
	BaseSHA        string     `gorm:"size:100"`
	Merged         bool
	Comments       int
	ReviewComments int `gorm:"column:review_comments"`
	Commits        int
	Additions      int
	Deletions      int
	ChangedFiles   int `gorm:"column:changed_files"`
}

type GitHubPullRequestEvent struct {
	ID            string `gorm:"type:uuid;primaryKey"`
	Action        string `gorm:"size:20;index"`
	Number        int
	PullRequestID string      `gorm:"type:uuid;index"`
	PullRequest   PullRequest `gorm:"foreignKey:PullRequestID;references:IdPullRequest"`
	RepositoryID  string      `gorm:"type:uuid;index"`
	Repository    Repository  `gorm:"foreignKey:RepositoryID;references:IdRepository"`
	SenderID      string      `gorm:"type:uuid;index"`
	Sender        GitHubUser  `gorm:"foreignKey:SenderID;references:IdGit"`
	CreatedAt     time.Time   `gorm:"autoCreateTime"`
	UpdatedAt     time.Time   `gorm:"autoUpdateTime"`
}

func (u *GitHubUser) BeforeCreate(tx *gorm.DB) error {
	if u.IdGit == "" {
		u.IdGit = uuid.New().String()
	}
	return nil
}

func (r *Repository) BeforeCreate(tx *gorm.DB) error {
	if r.IdRepository == "" {
		r.IdRepository = uuid.New().String()
	}
	return nil
}

func (p *PullRequest) BeforeCreate(tx *gorm.DB) error {
	if p.IdPullRequest == "" {
		p.IdPullRequest = uuid.New().String()
	}
	return nil
}

func (g *GitHubPullRequestEvent) BeforeCreate(tx *gorm.DB) error {
	if g.ID == "" {
		g.ID = uuid.New().String()
	}
	return nil
}

func init() {
	println("Initializing models...")
	// Ordem importante: primeiro as tabelas independentes, depois as dependentes
	RegisterModel(&GitHubUser{})
	RegisterModel(&Repository{})
	RegisterModel(&PullRequest{})
	RegisterModel(&GitHubPullRequestEvent{})
}
