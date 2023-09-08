package service_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cocoide/commitify/internal/service"
	mock_gateway "github.com/cocoide/commitify/mock"
	"github.com/golang/mock/gomock"
)

func Test_GenerateCommitMessage(t *testing.T) {
	stagingCode := "test"
	type test struct {
		ErrorCaseName  string
		ChatGptResult  string
		ModifiedResult []string
	}
	tests := []test{
		{"先頭にインデックスを含む", "1. test\n2. test", []string{"test", "test"}},
		{"先頭にハイフンを含む", "- test\n- test", []string{"test", "test"}},
		{"先頭にスペースを含む", " test\n test", []string{"test", "test"}},
		{"改行ができていない", "Add A function. Update B", []string{"Add A function", "Update B"}},
		// たまに,  改行せずピリオド『.』で文章を区切るパターンがある
	}
	ctrl := gomock.NewController(t)
	og := mock_gateway.NewMockOpenAIGateway(ctrl)
	ms := service.NewMessageService(og)

	for _, test := range tests {
		prompt := fmt.Sprintf(service.CommitMessagePrompt, stagingCode)
		og.EXPECT().
			GetAnswerFromPrompt(prompt, service.PromptVariability).
			Return(test.ChatGptResult, nil)

		serviceResult, err := ms.GenerateCommitMessage(stagingCode)
		if err != nil {
			t.Error(err.Error())
		}
		if !reflect.DeepEqual(test.ModifiedResult, serviceResult) {
			fmt.Printf("FAIL: %s\n", test.ErrorCaseName)
			t.Errorf("Exp: %v, Got: %v", test.ModifiedResult, serviceResult)
		}
	}
}
