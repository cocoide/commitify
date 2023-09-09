package entity

type Config struct {
	ChatGptApiKey string `json:"chatGptApiKey"`
	UseLanguage   string `json:"UseLanguage"`
	CommitFormat  string `json:"CommitFormat"`
}
