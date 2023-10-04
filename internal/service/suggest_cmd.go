package service

import (
	"log"

	"github.com/cocoide/commitify/internal/entity"
	"github.com/cocoide/commitify/internal/gateway"
)

type SuggestCmdService struct {
	ais gateway.AISourceGatewayInterface
	fds fileDiffService
}

func NewSuggestCmdService() (*SuggestCmdService, error) {
	conf, err := entity.ReadConfig()
	if err != nil {
		return nil, err
	}

	var aigi gateway.AISourceGatewayInterface
	switch conf.AISource {
	case int(entity.WrapServer):
		aigi = gateway.NewGrpcServeGateway()
	case int(entity.OpenAiAPI):
		log.Fatal("現在、非対応の機能です。")
		return nil, err
	default:
		aigi = gateway.NewGrpcServeGateway()
	}

	fds := NewFileDiffService()

	return &SuggestCmdService{ais: aigi, fds: fds}, nil
}

func (scs *SuggestCmdService) GenerateCommitMessages() ([]string, error) {
	fileDiffStr, err := scs.fds.createFileDiffStr()
	if err != nil {
		return nil, err
	}

	return scs.ais.FetchCommitMessages(fileDiffStr)
}
