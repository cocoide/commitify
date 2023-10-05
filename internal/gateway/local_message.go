package gateway

import (
	"fmt"
	"github.com/cocoide/commitify/internal/entity"
	"github.com/cocoide/commitify/internal/service"
	"regexp"
	"strings"
)

const (
	NormalMessagePrompt = "Generate up to 5 commit messages for [%s]. Each message should be separated by only space"
)

var CommitMessageRegex = regexp.MustCompile(`^(\d.\s+)|^(-\s+)|^(\s+)`)

type localMessageService struct {
	nlp service.NLPService
}

func NewLocalMessageService(nlp service.NLPService) service.CommitMessageService {
	return &localMessageService{nlp: nlp}
}

func (l *localMessageService) GenerateCommitMessageList(code string, conf entity.Config) ([]string, error) {
	prompt := fmt.Sprintf(NormalMessagePrompt, code)
	result, err := l.nlp.GetAnswerFromPrompt(prompt)
	if err != nil {
		return nil, err
	}
	messages := strings.Split(result, "\n")
	messages = l.removeFromArrayByRegex(messages, CommitMessageRegex)
	return messages, nil
}

func (l *localMessageService) removeFromArrayByRegex(array []string, pattern *regexp.Regexp) []string {
	for i, msg := range array {
		array[i] = pattern.ReplaceAllString(msg, "")
	}
	return array
}
