package service

import (
	"fmt"
	"strings"

	"github.com/cocoide/commitify/internal/gateway"
	"github.com/cocoide/commitify/util"
)

const (
	CommitMessagePrompt = "Generate up to 5 commit messages for [%s]. Each message should be separated by only space"
	FormatNotice        = ", format commit as:\n- feat: [feature description]\n- bugfix: [bugfix description]"
)

// メッセージの生成、加工に関するクラス
type MessageService interface {
	AsyncGenerateCommitMessage() ([]string, error)
}

type messageService struct {
	og gateway.OpenAIGateway
}

func NewMessageService(og gateway.OpenAIGateway) MessageService {
	return &messageService{og: og}
}

func (s *messageService) AsyncGenerateCommitMessage() ([]string, error) {
	var result <-chan string
	stagingCode := util.ExecGetStagingCode()
	if len(stagingCode) < 1 {
		return nil, fmt.Errorf("There is no staging code")
	}
	prompt := fmt.Sprintf(CommitMessagePrompt, string(stagingCode))
	result = s.og.AsyncGetAnswerFromPrompt(prompt, 0.01)
	messages := strings.Split(<-result, "\n")
	return messages, nil
}
