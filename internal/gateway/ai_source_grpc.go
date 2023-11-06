package gateway

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"github.com/cocoide/commitify/internal/entity"
	pb "github.com/cocoide/commitify/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	// serveAddress = "localhost:54322"
	serveAddress = "commitify.fly.dev:443"
)

type grpcServeGateway struct {
	client pb.CommitMessageServiceClient
}

func NewGrpcServeGateway() *grpcServeGateway {
	gsg := new(grpcServeGateway)

	conn, err := grpc.Dial(
		serveAddress,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
	)
	if err != nil {
		log.Printf("Connection failed: %v", err)
		os.Exit(-1)
		return nil
	}

	gsg.client = pb.NewCommitMessageServiceClient(conn)

	return gsg
}

func (gsg grpcServeGateway) FetchCommitMessages(fileDiffStr string) ([]string, error) {
	// 設定情報を取得
	conf, err := entity.ReadConfig()
	if err != nil {
		fmt.Printf("設定ファイルが開けません: %v", err)
	}
	cft, lt := conf.Config2PbVars()

	req := &pb.CommitMessageRequest{
		InputCode:  fileDiffStr,
		CodeFormat: cft,
		Language:   lt,
	}

	res, err := gsg.client.GenerateCommitMessage(context.Background(), req)
	if err != nil {
		log.Fatal("gRPCの送信に失敗: ", err)
		return nil, err
	}

	return res.GetMessages(), nil

}
