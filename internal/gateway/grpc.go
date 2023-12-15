package gateway

import (
	"crypto/tls"
	"github.com/cocoide/commitify/internal/entity"
	"github.com/cocoide/commitify/internal/service"
	pb "github.com/cocoide/commitify/proto/gen"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

const (
	serverAddress = "commitify.fly.dev:443"
)

type grpcServerGateway struct {
	client pb.CommitMessageServiceClient
}

func NewGrpcServerGateway() service.CommitMessageService {
	conn, err := grpc.Dial(
		serverAddress,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
	)
	if err != nil {
		log.Fatalf("Failed to setup grpc connection: %v", err)
	}
	client := pb.NewCommitMessageServiceClient(conn)
	return &grpcServerGateway{client: client}
}

func (g *grpcServerGateway) GenerateCommitMessageList(code string, conf entity.Config) ([]string, error) {
	formatType, languageType := conf.Config2PbVars()
	req := &pb.CommitMessageRequest{
		CodeFormat: formatType,
		Language:   languageType,
		InputCode:  code,
	}
	res, err := g.client.GenerateCommitMessage(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res.Messages, nil
}
