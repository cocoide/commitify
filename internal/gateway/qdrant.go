package gateway

import (
	"encoding/json"

	"github.com/cocoide/commitify/internal/entity"
	"github.com/cocoide/commitify/internal/service"
)

type qdrantServerGateway struct {
	client *HttpClient
}

func NewQdrantServerGateway() service.CommitMessageService {
	c := NewHttpClient()

	return &qdrantServerGateway{client: c}
}

func (qs *qdrantServerGateway) GenerateCommitMessageList(diff string, conf entity.Config) ([]string, error) {
	type qdrantBody struct {
		Diff string `json:"diff"`
	}

	body := &qdrantBody{
		Diff: diff,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	res, err := qs.client.
		WithBaseURL("http://suwageeks.org:5215").
		WithPath("/search").
		WithHeader("Content-Type", "application/json").
		WithBody(b).Execute(POST)
	if err != nil {
		return nil, err
	}

	type qdrantResponse struct {
		Messages []string `json:"messages"`
	}

	values := new(qdrantResponse)
	if err = json.Unmarshal(res, values); err != nil {
		return nil, err
	}

	return values.Messages, nil
}
