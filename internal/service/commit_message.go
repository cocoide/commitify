package service

import "github.com/cocoide/commitify/internal/entity"

//go:generate mockgen -source=commit_message.go -destination=../../mock/commit_message.go
type CommitMessageService interface {
	GenerateCommitMessageList(code string, config entity.Config) ([]string, error)
}
