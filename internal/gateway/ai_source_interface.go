package gateway

type AISourceGatewayInterface interface {
	FetchCommitMessages(fileDiffStr string) ([]string, error)
}
