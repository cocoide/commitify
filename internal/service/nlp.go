package service

type NLPService interface {
	GetAnswerFromPrompt(prompt string) (string, error)
}
