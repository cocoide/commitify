package gateway

import (
	"context"
	"crypto/tls"
	"log"
	"os"

	"github.com/cocoide/commitify/internal/service"
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

func (gsg grpcServeGateway) FetchCommitMessages() ([]string, error) {
	// 設定情報からOpenAIへのアクセス方法の変更
	// if conf, err := entity.ReadConfig(); err != nil {
	// 	fmt.Printf("設定ファイルが開けません: %v", err)
	// }

	fds := service.NewFileDiffService()

	diffStr, err := fds.CreateFileDiffStr()
	if err != nil {
		log.Fatal("差分の取得に失敗: ", err)
		return nil, err
	}

	req := &pb.CommitMessageRequest{
		InputCode:  diffStr,
		CodeFormat: pb.CodeFormatType_PREFIX,
		Language:   pb.LanguageType_ENGLISH,
	}

	res, err := gsg.client.GenerateCommitMessage(context.Background(), req)
	if err != nil {
		log.Fatal("gRPCの送信に失敗: ", err)
		return nil, err
	}

	return res.GetMessages(), nil

}
