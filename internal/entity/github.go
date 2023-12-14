package entity

type PullRequest struct {
	Owner string
	Repo  string
	Title string
	Body  string
	Head  string
	Base  string
}

type PullRequests []*PullRequest
