package gateway

import (
	"context"
	"log"

	"github.com/cocoide/commitify/util"
	"github.com/sashabaranov/go-openai"
)

//go:generate mockgen -source=openai.go -destination=../../mock/openai.go
type OpenAIGateway interface {
	GetAnswerFromPrompt(prompt string, variability float32) (string, error)
	AsyncGetAnswerFromPrompt(prompt string, variability float32) <-chan string
}

type openAIGateway struct {
	client *openai.Client
	ctx    context.Context
}

func NewOpenAIGateway(ctx context.Context) OpenAIGateway {
	config, err := util.ReadConfig()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}
	client := openai.NewClient(config.ChatGptApiKey)
	return &openAIGateway{client: client, ctx: ctx}
}

func (og *openAIGateway) GetAnswerFromPrompt(prompt string, variability float32) (string, error) {
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		Temperature: variability,
	}
	res, err := og.client.CreateChatCompletion(og.ctx, req)
	if err != nil {
		return "", err
	}
	answer := res.Choices[0].Message.Content
	return answer, nil
}

func (og *openAIGateway) AsyncGetAnswerFromPrompt(prompt string, variability float32) <-chan string {
	responseCh := make(chan string, 1)

	go func() {
		answer, _ := og.GetAnswerFromPrompt(prompt, variability)
		responseCh <- answer
	}()

	return responseCh
}
