package usecase

import (
	"fmt"
	"github.com/cocoide/commitify/internal/entity"
	"github.com/cocoide/commitify/internal/service"
	"github.com/pkg/errors"
	"strings"
)

type PushCmdUsecase struct {
	github service.GithubService
	nlp    service.NLPService
}

const (
	GeneratePRFromCommitsPrompt = "generate pull request content from commit messages [%s] in Japanese"
	GeneratePRTitleFromPRBody   = "generate pull request title from pr body [%s] in Japanese as one sentence like 〇〇の機能追加"
)

func NewPushCmdUsecase(github service.GithubService, nlp service.NLPService) *PushCmdUsecase {
	return &PushCmdUsecase{github: github, nlp: nlp}
}

func (u *PushCmdUsecase) GetRemoteBaseBranchCandidates() ([]string, error) {
	return u.github.GetRecentUpdatedBranch()
}

func (u *PushCmdUsecase) GeneratePullRequest(base string) (*entity.PullRequest, error) {
	head, err := u.github.GetCurrentBranch()
	if err != nil {
		return nil, err
	}
	details, err := u.github.GetCurrentRepoDetails()
	if err != nil {
		return nil, err
	}
	commits, err := u.github.GetUnPushedCommits(base)
	if err != nil {
		return nil, err
	}
	prBodyPrompt := fmt.Sprintf(GeneratePRFromCommitsPrompt, strings.Join(commits, ", "))
	body, err := u.nlp.GetAnswerFromPrompt(prBodyPrompt)
	if err != nil {
		return nil, err
	}
	prTitlePrompt := fmt.Sprintf(GeneratePRTitleFromPRBody, body)
	title, err := u.nlp.GetAnswerFromPrompt(prTitlePrompt)
	if err != nil {
		return nil, err
	}
	return &entity.PullRequest{
		Head:  head,
		Base:  base,
		Body:  body,
		Title: title,
		Repo:  details.Repo,
		Owner: details.Owner,
	}, nil
}

func (u *PushCmdUsecase) SubmitPullRequest(pr *entity.PullRequest) error {
	config, err := entity.ReadConfig()
	if err != nil {
		return err
	}
	tkn := config.GithubToken
	if tkn == "" {
		return errors.New("github token not found")
	}
	return u.github.CreatePullRequest(pr, tkn)
}
