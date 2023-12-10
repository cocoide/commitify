package service

import "github.com/cocoide/commitify/internal/entity"

// 分割コミットの生成のクライアント側もここに入れていく
type CommitMessageService interface {
	GenerateCommitMessageList(code string, config entity.Config) ([]string, error)
}

// githubに関するデータのinput/output
type GithubService interface {
	GetStagingCodeDiff() (string, error)
}

type NLPService interface {
	GetAnswerFromPrompt(prompt string) (string, error)
}
