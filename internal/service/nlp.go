package service

//go:generate mockgen -source=nlp.go -destination=../../mock/nlp.go
type NLPService interface {
	GetAnswerFromPrompt(prompt string) (string, error)
}
