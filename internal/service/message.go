package service

// import (
// 	"fmt"
// 	"strings"

// 	"github.com/cocoide/commitify/internal/gateway"
// )

// const (
// 	CommitMessagePrompt = "Generate up to 5 commit messages for [%s]. Each message should be separated by only space"
// 	FormatNotice        = ", format commit as:\n- feat: [feature description]\n- bugfix: [bugfix description]"
// )

// var PromptVariability float32 = 0.01

// // メッセージの生成、加工に関するクラス
// type MessageService interface {
// 	GenerateCommitMessage(stagingCode string) ([]string, error)
// }

// type messageService struct {
// 	og gateway.OpenAIGateway
// }

// func NewMessageService(og gateway.OpenAIGateway) MessageService {
// 	return &messageService{og: og}
// }

// func (s *messageService) GenerateCommitMessage(stagingCode string) ([]string, error) {
// 	if len(stagingCode) < 1 {
// 		return nil, fmt.Errorf("There is no staging code")
// 	}
// 	prompt := fmt.Sprintf(CommitMessagePrompt, stagingCode)
// 	result, err := s.og.GetAnswerFromPrompt(prompt, PromptVariability)
// 	if err != nil {
// 		return nil, err
// 	}
// 	messages := strings.Split(result, "\n")
// 	return messages, nil
// }
