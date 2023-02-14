package clients

import (
	"github.com/oliver7100/token-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewTokenClient(cfg *clientConfig) (proto.AuthServiceClient, error) {
	d, err := grpc.Dial(cfg.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return proto.NewAuthServiceClient(d), nil
}
