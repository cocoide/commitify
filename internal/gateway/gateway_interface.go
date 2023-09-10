package gateway

type GatewayInterface interface {
	FetchCommitMessages() ([]string, error)
}
