package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type GitHubUser struct {
	IdGit     string `gorm:"column:id_git;type:uuid;primaryKey"`
	ID        int    `gorm:"column:id;index" json:"id"`
	NodeID    string `gorm:"column:node_id;size:50" json:"node_id"`
	Login     string `gorm:"size:100" json:"login"`
	AvatarURL string `gorm:"column:avatar_url;size:255" json:"avatar_url"`
	HTMLURL   string `gorm:"column:html_url;size:255" json:"html_url"`
	Type      string `gorm:"size:50" json:"type"`
	SiteAdmin bool   `gorm:"column:site_admin" json:"site_admin"`
}

type Link struct {
	BaseModel
	Href string `gorm:"size:255" json:"href"`
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
	OwnerID        string     `gorm:"type:uuid;index"`
	Owner          GitHubUser `gorm:"foreignKey:OwnerID;references:IdGit"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	PushedAt       time.Time `gorm:"column:pushed_at"`
}

type Branch struct {
	BaseModel
	Label  string      `gorm:"size:255" json:"label"`
	Ref    string      `gorm:"size:255" json:"ref"`
	Sha    string      `gorm:"size:100" json:"sha"`
	UserID string      `gorm:"type:uuid;index" json:"-"`
	User   *GitHubUser `gorm:"foreignKey:UserID;references:IdGit;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	RepoID string      `gorm:"type:uuid;index" json:"-"`
	Repo   *Repository `gorm:"foreignKey:RepoID;references:IdRepository;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"repo"`
}

type PRLinks struct {
	BaseModel
	PullRequestID    string `gorm:"type:uuid;index" json:"-"`
	SelfID           string `gorm:"type:uuid;index" json:"-"`
	Self             Link   `gorm:"foreignKey:SelfID;references:ID" json:"self"`
	HTMLID           string `gorm:"type:uuid;index" json:"-"`
	HTML             Link   `gorm:"foreignKey:HTMLID;references:ID" json:"html"`
	IssueID          string `gorm:"type:uuid;index" json:"-"`
	Issue            Link   `gorm:"foreignKey:IssueID;references:ID" json:"issue"`
	CommentsID       string `gorm:"type:uuid;index" json:"-"`
	Comments         Link   `gorm:"foreignKey:CommentsID;references:ID" json:"comments"`
	ReviewCommentsID string `gorm:"type:uuid;index" json:"-"`
	ReviewComments   Link   `gorm:"foreignKey:ReviewCommentsID;references:ID" json:"review_comments"`
	ReviewCommentID  string `gorm:"type:uuid;index" json:"-"`
	ReviewComment    Link   `gorm:"foreignKey:ReviewCommentID;references:ID" json:"review_comment"`
	CommitsID        string `gorm:"type:uuid;index" json:"-"`
	Commits          Link   `gorm:"foreignKey:CommitsID;references:ID" json:"commits"`
	StatusesID       string `gorm:"type:uuid;index" json:"-"`
	Statuses         Link   `gorm:"foreignKey:StatusesID;references:ID" json:"statuses"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

type PullRequest struct {
	IdPullRequest       string      `gorm:"column:id_pull_request;type:uuid;primaryKey"`
	URL                 string      `gorm:"size:255" json:"url"`
	ID                  int         `gorm:"column:id;index" json:"id"`
	NodeID              string      `gorm:"size:50" json:"node_id"`
	HTMLURL             string      `gorm:"size:255" json:"html_url"`
	DiffURL             string      `gorm:"size:255" json:"diff_url"`
	PatchURL            string      `gorm:"size:255" json:"patch_url"`
	IssueURL            string      `gorm:"size:255" json:"issue_url"`
	Number              int         `gorm:"index" json:"number"`
	State               string      `gorm:"size:50" json:"state"`
	Locked              bool        `json:"locked"`
	Title               string      `gorm:"size:255" json:"title"`
	UserID              *string     `gorm:"type:uuid;index" json:"-"`
	User                *GitHubUser `gorm:"foreignKey:UserID;references:IdGit;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	Body                *string     `gorm:"type:text" json:"body"`
	CreatedAt           time.Time   `json:"created_at"`
	UpdatedAt           time.Time   `json:"updated_at"`
	ClosedAt            *time.Time  `json:"closed_at"`
	MergedAt            *time.Time  `json:"merged_at"`
	MergeCommitSHA      string      `gorm:"column:merge_commit_sha;size:100" json:"merge_commit_sha"`
	AssigneeID          *string     `gorm:"type:uuid;index" json:"-"`
	Assignee            *GitHubUser `gorm:"foreignKey:AssigneeID;references:IdGit;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"assignee"`
	Draft               bool        `json:"draft"`
	HeadID              *string     `gorm:"type:uuid;index" json:"-"`
	Head                *Branch     `gorm:"foreignKey:HeadID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"head"`
	BaseID              *string     `gorm:"type:uuid;index" json:"-"`
	Base                *Branch     `gorm:"foreignKey:BaseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"base"`
	AuthorAssociation   string      `gorm:"size:100" json:"author_association"`
	Merged              bool        `json:"merged"`
	Mergeable           *bool       `json:"mergeable"`
	Rebaseable          *bool       `json:"rebaseable"`
	MergeableState      string      `gorm:"size:50" json:"mergeable_state"`
	Comments            int         `json:"comments"`
	ReviewComments      int         `gorm:"column:review_comments" json:"review_comments"`
	MaintainerCanModify bool        `gorm:"column:maintainer_can_modify" json:"maintainer_can_modify"`
	Commits             int         `json:"commits"`
	Additions           int         `json:"additions"`
	Deletions           int         `json:"deletions"`
	ChangedFiles        int         `gorm:"column:changed_files" json:"changed_files"`
}

type GitHubEvent struct {
	BaseModel
	Action        string       `gorm:"size:50" json:"action"`
	Number        int          `gorm:"index" json:"number"`
	PullRequestID string       `gorm:"index" json:"-"`
	PullRequest   *PullRequest `gorm:"foreignKey:PullRequestID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"pull_request"`
	RepositoryID  string       `gorm:"index" json:"-"`
	Repository    *Repository  `gorm:"foreignKey:RepositoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"repository"`
	SenderID      string       `gorm:"index" json:"-"`
	Sender        *GitHubUser  `gorm:"foreignKey:SenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"sender"`
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

func (b *Branch) BeforeCreate(tx *gorm.DB) error {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return nil
}

func (l *Link) BeforeCreate(tx *gorm.DB) error {
	if l.ID == "" {
		l.ID = uuid.New().String()
	}
	return nil
}

func (pr *PullRequest) BeforeCreate(tx *gorm.DB) error {
	if pr.IdPullRequest == "" {
		pr.IdPullRequest = uuid.New().String()
	}

	return nil
}

func (e *GitHubEvent) BeforeCreate(tx *gorm.DB) error {
	if e.ID == "" {
		e.ID = uuid.New().String()
	}
	return nil
}

func init() {
	println("Initializing models...")
	RegisterModel(&GitHubUser{})
	RegisterModel(&Link{})
	RegisterModel(&Repository{})
	RegisterModel(&Branch{})
	RegisterModel(&PullRequest{})
	RegisterModel(&GitHubEvent{})
}
