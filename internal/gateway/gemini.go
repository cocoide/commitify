package gateway

import (
	"encoding/json"

	"github.com/cocoide/commitify/internal/entity"
	"github.com/cocoide/commitify/internal/service"
	"github.com/pkg/errors"
)

type geminiServerGateway struct {
	client *HttpClient
}

func NewGeminiServerGateway() service.CommitMessageService {
	c := NewHttpClient().
		WithBaseURL("http://suwageeks.org:5215").
		WithPath("/gemini").
		WithHeader("Content-Type", "application/json")

	return &geminiServerGateway{client: c}
}

func (qs *geminiServerGateway) GenerateCommitMessageList(diff string, conf entity.Config) ([]string, error) {
	if diff == "" {
		return nil, errors.New("ステージされた変更がありません。")
	}

	type geminiBody struct {
		Diff string `json:"diff"`
	}

	body := &geminiBody{
		Diff: diff,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	res, err := qs.client.WithBody(b).Execute(POST)
	if err != nil {
		return nil, err
	}

	type geminiResponse struct {
		Messages []string `json:"messages"`
	}

	values := new(geminiResponse)
	if err = json.Unmarshal(res, values); err != nil {
		return nil, err
	}

	return values.Messages, nil
}
