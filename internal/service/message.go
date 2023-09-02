package service

import (
	"fmt"
	"os/exec"

	"github.com/cocoide/commitify/internal/gateway"
)

const (
	CommitMessagePrompt = "Generate commit message for [%s]"
	FormatNotice        = ", format commit as:\n- feat: [feature description]\n- bugfix: [bugfix description]"
)

// メッセージの生成、加工に関するクラス
type MessageService interface {
	AsyncGenerateCommitMessage() (<-chan string, error)
}

type messageService struct {
	og gateway.OpenAIGateway
}

func NewMessageService(og gateway.OpenAIGateway) MessageService {
	return &messageService{og: og}
}

func (s *messageService) AsyncGenerateCommitMessage() (<-chan string, error) {
	var result <-chan string
	stagingCode, err := exec.Command("git", "diff", "--staged").Output()
	if err != nil {
		return nil, err
	}
	prompt := fmt.Sprintf(CommitMessagePrompt, string(stagingCode))
	result = s.og.AsyncGetAnswerFromPrompt(prompt, 0.01)
	return result, nil
}
