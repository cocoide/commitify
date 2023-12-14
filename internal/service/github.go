package service

import "github.com/cocoide/commitify/internal/entity"

type GetRepoDetailsResponse struct {
	Owner string
	Repo  string
}

type GithubService interface {
	GetStagingCodeDiff() (string, error)
	GetCurrentRepoDetails() (*GetRepoDetailsResponse, error)
	CreatePullRequest(pr *entity.PullRequest, token string) error
	GetCurrentBranch() (string, error)
	GetUnPushedCommits(base string) ([]string, error)
	GetRecentUpdatedBranch() ([]string, error)
}
