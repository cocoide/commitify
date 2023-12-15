package gateway

import (
	"encoding/json"
	"fmt"
	"github.com/cocoide/commitify/internal/entity"
	"github.com/cocoide/commitify/internal/service"
	"github.com/pkg/errors"
	"os/exec"
	"strings"
)

type githubGateway struct {
}

func NewGithubGateway() service.GithubService {
	return &githubGateway{}
}

func (g *githubGateway) GetStagingCodeDiff() (string, error) {
	// Gitが入ってるかどうかのチェックも入れる
	// 入っていないなら専用のエラーメッセージを生成
	diff, err := exec.Command("git", "diff", "--staged").Output()
	return string(diff), err
}

func (g *githubGateway) GetCurrentRepoDetails() (*service.GetRepoDetailsResponse, error) {
	_, err := exec.LookPath("git")
	if err != nil {
		return nil, errors.New("git is not installed on the system")
	}
	output, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()
	if err != nil {
		return nil, err
	}
	url := strings.TrimSpace(string(output))
	parts := strings.Split(url, "/")
	if len(parts) < 2 {
		return nil, errors.New("unable to parse the repository URL")
	}
	repo := strings.TrimSuffix(parts[len(parts)-1], ".git")
	owner := parts[len(parts)-2]

	return &service.GetRepoDetailsResponse{
		Owner: owner,
		Repo:  repo,
	}, nil
}

func (g *githubGateway) CreatePullRequest(req *entity.PullRequest, token string) error {
	type Body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
		Head  string `json:"head"`
		Base  string `json:"base"`
	}
	body := &Body{
		Title: req.Title,
		Body:  req.Body,
		Head:  req.Head,
		Base:  req.Base,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	_, err = NewHttpClient().
		WithBaseURL(fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls", req.Owner, req.Repo)).
		WithHeader("Accept", "application/vnd.github+json").
		WithHeader("X-GitHub-Api-Version", "2022-11-28").
		WithBearerToken(token).
		WithBody(b).
		Execute(POST)
	return err
}

func (g *githubGateway) GetCurrentBranch() (string, error) {
	output, err := exec.Command("git", "branch", "--show-current").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(strings.TrimSpace(string(output))), nil
}

func (g *githubGateway) GetUnPushedCommits(base string) ([]string, error) {
	head, err := g.GetCurrentBranch()
	if err != nil {
		return nil, err
	}
	output, err := exec.Command("git", "log", base+".."+head, "--pretty=format:%s").Output()
	if err != nil {
		return nil, err
	}
	commits := strings.Split(string(output), "\n")
	return commits, nil
}

func (g *githubGateway) GetRecentUpdatedBranch() ([]string, error) {
	var result []string
	output, err := exec.Command("git", "for-each-ref", "--sort=-committerdate", "refs/remotes/", "--format=%(refname:short)").Output()
	if err != nil {
		return nil, err
	}
	for _, line := range strings.Split(string(output), "\n") {
		branch := strings.TrimPrefix(line, "origin/")
		if branch == "origin" {
			continue
		}
		result = append(result, branch)
	}
	return result, nil
}
