package service

import "github.com/cocoide/commitify/internal/entity"

type CommitMessageService interface {
	GenerateCommitMessageList(code string, config entity.Config) ([]string, error)
}
