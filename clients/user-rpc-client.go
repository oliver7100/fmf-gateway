package clients

import (
	"github.com/oliver7100/user-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserServiceClient(cfg *clientConfig) (proto.UserServiceClient, error) {
	c, err := grpc.Dial(cfg.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return proto.NewUserServiceClient(c), nil
}
