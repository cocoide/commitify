package gateway

import (
	"context"
	"github.com/cocoide/commitify/internal/service"
	"log"

	"github.com/cocoide/commitify/internal/entity"
	"github.com/sashabaranov/go-openai"
)

type openAIGateway struct {
	client *openai.Client
	ctx    context.Context
}

func NewOpenAIGateway(ctx context.Context) service.NLPService {
	config, err := entity.ReadConfig()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}
	client := openai.NewClient(config.ChatGptApiKey)
	return &openAIGateway{client: client, ctx: ctx}
}

func (og *openAIGateway) GetAnswerFromPrompt(prompt string) (string, error) {
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		Temperature: 0.001,
	}
	res, err := og.client.CreateChatCompletion(og.ctx, req)
	if err != nil {
		return "", err
	}
	answer := res.Choices[0].Message.Content
	return answer, nil
}
