package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type GitHubUser struct {
	IdGit     string `gorm:"column:id_git;type:uuid;primaryKey" `
	ID        int    `gorm:"column:id;index" json:"id"`
	NodeID    string `gorm:"column:node_id;size:50" json:"node_id"`
	Login     string `gorm:"size:100" json:"login"`
	AvatarURL string `gorm:"column:avatar_url;size:255" json:"avatar_url"`
	HTMLURL   string `gorm:"column:html_url;size:255" json:"html_url"`
	Type      string `gorm:"size:50" json:"type"`
	SiteAdmin bool   `gorm:"column:site_admin" json:"site_admin"`
}

type Repository struct {
	IdRepository   string     `gorm:"column:id_repository;type:uuid;primaryKey"`
	ID             int        `gorm:"column:id;index" json:"id"`
	NodeID         string     `gorm:"column:node_id;size:50" json:"node_id"`
	Name           string     `gorm:"size:100" json:"name"`
	FullName       string     `gorm:"column:full_name;size:200" json:"full_name"`
	Private        bool       `gorm:"column:private" json:"private"`
	HTMLURL        string     `gorm:"column:html_url;size:255" json:"html_url"`
	Description    string     `gorm:"type:text" json:"description"`
	Fork           bool       `gorm:"column:fork" json:"fork"`
	Language       string     `gorm:"size:50" json:"language"`
	HasIssues      bool       `gorm:"column:has_issues" json:"has_issues"`
	HasProjects    bool       `gorm:"column:has_projects" json:"has_projects"`
	HasWiki        bool       `gorm:"column:has_wiki" json:"has_wiki"`
	HasPages       bool       `gorm:"column:has_pages" json:"has_pages"`
	HasDiscussions bool       `gorm:"column:has_discussions" json:"has_discussions"`
	Archived       bool       `gorm:"column:archived" json:"archived"`
	Disabled       bool       `gorm:"column:disabled" json:"disabled"`
	OwnerID        string     `gorm:"type:uuid" `
	Owner          GitHubUser `gorm:"foreignKey:OwnerID;references:IdGit"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	PushedAt       time.Time `gorm:"column:pushed_at"`
}

type PullRequest struct {
	IdPullRequest  string     `gorm:"column:id_pull_request;type:uuid;primaryKey"`
	ID             int        `gorm:"column:id;index" json:"id"`
	NodeID         string     `gorm:"column:node_id;size:50" json:"node_id"`
	Number         int        `gorm:"index" json:"number"`
	State          string     `gorm:"size:50" json:"state"`
	Locked         bool       `json:"locked"`
	Title          string     `gorm:"size:255" json:"title"`
	Body           *string    `gorm:"type:text" json:"body"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	ClosedAt       *time.Time `json:"closed_at"`
	MergedAt       *time.Time `json:"merged_at"`
	MergeCommitSHA *string    `gorm:"column:merge_commit_sha;size:100" json:"merge_commit_sha"`
	Draft          bool       `json:"draft"`
	UserID         string     `gorm:"type:uuid"`
	User           GitHubUser `gorm:"foreignKey:UserID;references:IdGit" json:"user"`
	HTMLURL        string     `gorm:"column:html_url;size:255" json:"html_url"`
	HeadRef        string     `gorm:"size:100" json:"head_ref"`
	HeadSHA        string     `gorm:"size:100" json:"head_sha"`
	BaseRef        string     `gorm:"size:100" json:"base_ref"`
	BaseSHA        string     `gorm:"size:100" json:"base_sha"`
	Merged         bool       `json:"merged"`
	Comments       int        `json:"comments"`
	ReviewComments int        `json:"review_comments"`
	Commits        int        `json:"commits"`
	Additions      int        `json:"additions"`
	Deletions      int        `json:"deletions"`
	ChangedFiles   int        `json:"changed_files"`
}

type GitHubPullRequestEvent struct {
	ID            string      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Action        string      `gorm:"size:20;index" json:"action"`
	Number        int         `gorm:"index" json:"number"`
	PullRequestID string      `gorm:"type:uuid;index"`
	PullRequest   PullRequest `gorm:"foreignKey:PullRequestID;references:IdPullRequest;save_associations:true" json:"pull_request"`
	RepositoryID  string      `gorm:"type:uuid;index"`
	Repository    Repository  `gorm:"foreignKey:RepositoryID;references:IdRepository;save_associations:true" json:"repository"`
	SenderID      string      `gorm:"type:uuid;index"`
	Sender        GitHubUser  `gorm:"foreignKey:SenderID;references:IdGit;save_associations:true" json:"sender"`
	CreatedAt     time.Time   `gorm:"autoCreateTime"`
	UpdatedAt     time.Time   `gorm:"autoUpdateTime"`
}

func (g *GitHubPullRequestEvent) BeforeCreate(tx *gorm.DB) error {
	if g.ID == "" {
		g.ID = uuid.New().String()
	}
	if g.Sender.IdGit == "" {
		g.Sender.IdGit = uuid.New().String()
		g.SenderID = g.Sender.IdGit
	}
	if g.Repository.IdRepository == "" {
		g.Repository.IdRepository = uuid.New().String()
		g.RepositoryID = g.Repository.IdRepository
	}
	// Garantir que os IDs sejam gerados para as associações
	if g.PullRequest.IdPullRequest == "" {
		g.PullRequest.IdPullRequest = uuid.New().String()
		g.PullRequestID = g.PullRequest.IdPullRequest
	}

	if g.Repository.IdRepository == "" {
		g.Repository.IdRepository = uuid.New().String()
		g.RepositoryID = g.Repository.IdRepository
	}

	if g.Sender.IdGit == "" {
		g.Sender.IdGit = uuid.New().String()
		g.SenderID = g.Sender.IdGit
	}
	if g.PullRequest.User.IdGit == "" && g.PullRequest.User.ID != 0 {
		g.PullRequest.User.IdGit = uuid.New().String()
		g.PullRequest.UserID = g.PullRequest.User.IdGit
	}
	if g.Repository.Owner.IdGit == "" && g.Repository.Owner.ID != 0 {
		g.Repository.Owner.IdGit = uuid.New().String()
		g.Repository.OwnerID = g.Repository.Owner.IdGit
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
