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
	BaseModel
	IdGit             string `gorm:"column:id_git;type:uuid;primaryKey"`
	ID                int    `gorm:"column:id;index" json:"id"`
	NodeID            string `gorm:"column:node_id;size:50" json:"node_id"`
	Login             string `gorm:"size:100" json:"login"`
	AvatarURL         string `gorm:"column:avatar_url;size:255" json:"avatar_url"`
	GravatarID        string `gorm:"column:gravatar_id;size:255" json:"gravatar_id"`
	URL               string `gorm:"column:url;size:255" json:"url"`
	HTMLURL           string `gorm:"column:html_url;size:255" json:"html_url"`
	FollowersURL      string `gorm:"column:followers_url;size:255" json:"followers_url"`
	FollowingURL      string `gorm:"column:following_url;size:255" json:"following_url"`
	GistsURL          string `gorm:"column:gists_url;size:255" json:"gists_url"`
	StarredURL        string `gorm:"column:starred_url;size:255" json:"starred_url"`
	SubscriptionsURL  string `gorm:"column:subscriptions_url;size:255" json:"subscriptions_url"`
	OrganizationsURL  string `gorm:"column:organizations_url;size:255" json:"organizations_url"`
	ReposURL          string `gorm:"column:repos_url;size:255" json:"repos_url"`
	EventsURL         string `gorm:"column:events_url;size:255" json:"events_url"`
	ReceivedEventsURL string `gorm:"column:received_events_url;size:255" json:"received_events_url"`
	Type              string `gorm:"size:50" json:"type"`
	UserViewType      string `gorm:"column:user_view_type;size:50" json:"user_view_type"`
	SiteAdmin         bool   `gorm:"column:site_admin" json:"site_admin"`
}

type Link struct {
	BaseModel
	Href string `gorm:"size:255" json:"href"`
}

type Repository struct {
	BaseModel
	ID                       int         `gorm:"index" json:"id"`
	NodeID                   string      `gorm:"size:50" json:"node_id"`
	Name                     string      `gorm:"size:255" json:"name"`
	FullName                 string      `gorm:"column:full_name;size:255" json:"full_name"`
	Private                  bool        `json:"private"`
	OwnerID                  uint        `gorm:"index" json:"-"`
	Owner                    *GitHubUser `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"owner"`
	HTMLURL                  string      `gorm:"column:html_url;size:255" json:"html_url"`
	Description              string      `gorm:"type:text" json:"description"`
	Fork                     bool        `json:"fork"`
	URL                      string      `gorm:"size:255" json:"url"`
	CreatedAt                time.Time   `json:"created_at"`
	UpdatedAt                time.Time   `json:"updated_at"`
	PushedAt                 time.Time   `json:"pushed_at"`
	Homepage                 string      `gorm:"size:255" json:"homepage"`
	Size                     int         `json:"size"`
	StargazersCount          int         `gorm:"column:stargazers_count" json:"stargazers_count"`
	WatchersCount            int         `gorm:"column:watchers_count" json:"watchers_count"`
	Language                 string      `gorm:"size:100" json:"language"`
	HasIssues                bool        `gorm:"column:has_issues" json:"has_issues"`
	HasProjects              bool        `gorm:"column:has_projects" json:"has_projects"`
	HasDownloads             bool        `gorm:"column:has_downloads" json:"has_downloads"`
	HasWiki                  bool        `gorm:"column:has_wiki" json:"has_wiki"`
	HasPages                 bool        `gorm:"column:has_pages" json:"has_pages"`
	HasDiscussions           bool        `gorm:"column:has_discussions" json:"has_discussions"`
	ForksCount               int         `gorm:"column:forks_count" json:"forks_count"`
	Archived                 bool        `json:"archived"`
	Disabled                 bool        `json:"disabled"`
	OpenIssuesCount          int         `gorm:"column:open_issues_count" json:"open_issues_count"`
	AllowForking             bool        `gorm:"column:allow_forking" json:"allow_forking"`
	IsTemplate               bool        `gorm:"column:is_template" json:"is_template"`
	WebCommitSignoffRequired bool        `gorm:"column:web_commit_signoff_required" json:"web_commit_signoff_required"`
	Topics                   []string    `gorm:"serializer:json" json:"topics"`
	Visibility               string      `gorm:"size:50" json:"visibility"`
	Forks                    int         `json:"forks"`
	OpenIssues               int         `gorm:"column:open_issues" json:"open_issues"`
	Watchers                 int         `json:"watchers"`
	DefaultBranch            string      `gorm:"column:default_branch;size:100" json:"default_branch"`
}

type Branch struct {
	BaseModel
	Label  string      `gorm:"size:255" json:"label"`
	Ref    string      `gorm:"size:255" json:"ref"`
	Sha    string      `gorm:"size:100" json:"sha"`
	UserID uint        `gorm:"index" json:"-"`
	User   *GitHubUser `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	RepoID uint        `gorm:"index" json:"-"`
	Repo   *Repository `gorm:"foreignKey:RepoID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"repo"`
}

type PRLinks struct {
	BaseModel
	PullRequestID    uint `gorm:"index" json:"-"`
	SelfID           uint `gorm:"index" json:"-"`
	Self             Link `gorm:"foreignKey:SelfID" json:"self"`
	HTMLID           uint `gorm:"index" json:"-"`
	HTML             Link `gorm:"foreignKey:HTMLID" json:"html"`
	IssueID          uint `gorm:"index" json:"-"`
	Issue            Link `gorm:"foreignKey:IssueID" json:"issue"`
	CommentsID       uint `gorm:"index" json:"-"`
	Comments         Link `gorm:"foreignKey:CommentsID" json:"comments"`
	ReviewCommentsID uint `gorm:"index" json:"-"`
	ReviewComments   Link `gorm:"foreignKey:ReviewCommentsID" json:"review_comments"`
	ReviewCommentID  uint `gorm:"index" json:"-"`
	ReviewComment    Link `gorm:"foreignKey:ReviewCommentID" json:"review_comment"`
	CommitsID        uint `gorm:"index" json:"-"`
	Commits          Link `gorm:"foreignKey:CommitsID" json:"commits"`
	StatusesID       uint `gorm:"index" json:"-"`
	Statuses         Link `gorm:"foreignKey:StatusesID" json:"statuses"`
}

type PullRequest struct {
	BaseModel
	URL                 string      `gorm:"size:255" json:"url"`
	ID                  int         `gorm:"index" json:"id"`
	NodeID              string      `gorm:"size:50" json:"node_id"`
	HTMLURL             string      `gorm:"size:255" json:"html_url"`
	DiffURL             string      `gorm:"size:255" json:"diff_url"`
	PatchURL            string      `gorm:"size:255" json:"patch_url"`
	IssueURL            string      `gorm:"size:255" json:"issue_url"`
	Number              int         `gorm:"index" json:"number"`
	State               string      `gorm:"size:50" json:"state"`
	Locked              bool        `json:"locked"`
	Title               string      `gorm:"size:255" json:"title"`
	UserID              uint        `gorm:"index" json:"-"`
	User                *GitHubUser `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	Body                *string     `gorm:"type:text" json:"body"`
	CreatedAt           time.Time   `json:"created_at"`
	UpdatedAt           time.Time   `json:"updated_at"`
	ClosedAt            *time.Time  `json:"closed_at"`
	MergedAt            *time.Time  `json:"merged_at"`
	MergeCommitSHA      string      `gorm:"column:merge_commit_sha;size:100" json:"merge_commit_sha"`
	AssigneeID          uint        `gorm:"index" json:"-"`
	Assignee            *GitHubUser `gorm:"foreignKey:AssigneeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"assignee"`
	Draft               bool        `json:"draft"`
	HeadID              uint        `gorm:"index" json:"-"`
	Head                *Branch     `gorm:"foreignKey:HeadID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"head"`
	BaseID              uint        `gorm:"index" json:"-"`
	Base                *Branch     `gorm:"foreignKey:BaseID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"base"`
	LinksID             uint        `gorm:"index" json:"-"`
	Links               *PRLinks    `gorm:"foreignKey:LinksID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"_links"`
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
	PullRequestID uint         `gorm:"index" json:"-"`
	PullRequest   *PullRequest `gorm:"foreignKey:PullRequestID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"pull_request"`
	RepositoryID  uint         `gorm:"index" json:"-"`
	Repository    *Repository  `gorm:"foreignKey:RepositoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"repository"`
	SenderID      uint         `gorm:"index" json:"-"`
	Sender        *GitHubUser  `gorm:"foreignKey:SenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"sender"`
}

func (g *GitHubEvent) BeforeCreate(tx *gorm.DB) error {
	println("BeforeCreate")
	if tx.Statement.Schema.PrioritizedPrimaryField.DBName == "id" {
		if id := tx.Statement.Schema.PrioritizedPrimaryField.ReflectValueOf(tx.Statement.Context, tx.Statement.ReflectValue); id.IsZero() {
			tx.Statement.SetColumn("id", uuid.New().String())
		}
	}
	return nil
}

func init() {
	println("Initializing models...")
	RegisterModel(&GitHubUser{})
	RegisterModel(&Link{})
	RegisterModel(&Repository{})
	RegisterModel(&Branch{})
	RegisterModel(&PRLinks{})
	RegisterModel(&PullRequest{})
	RegisterModel(&GitHubEvent{})
}
