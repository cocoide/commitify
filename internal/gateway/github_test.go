package gateway

import (
	"github.com/cocoide/commitify/internal/entity"
	"log"
	"testing"
)

const (
	TestHeadBranch = "temp/test-auto-pr"
)

func Test_CreatePullRequest(t *testing.T) {
	conf, err := entity.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	tkn := conf.GithubToken
	pr := &entity.PullRequest{
		Owner: "cocoide", Repo: "commitify", Title: "test title", Body: "test body", Head: TestHeadBranch, Base: "main"}
	u := NewGithubGateway()
	if err := u.CreatePullRequest(pr, tkn); err != nil {
		t.Error(err)
	}
}
