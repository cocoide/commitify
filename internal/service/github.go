package service

import "github.com/cocoide/commitify/internal/entity"

type GetRepoDetailsResponse struct {
	Owner string
	Repo  string
}

//go:generate mockgen -source=github.go -destination=../../mock/github.go
type GithubService interface {
	GetStagingCodeDiff() (string, error)
	GetCurrentRepoDetails() (*GetRepoDetailsResponse, error)
	CreatePullRequest(pr *entity.PullRequest, token string) error
	GetCurrentBranch() (string, error)
	GetUnPushedCommits(base string) ([]string, error)
	GetRecentUpdatedBranch() ([]string, error)
	PushCurrentBranch() error
}
