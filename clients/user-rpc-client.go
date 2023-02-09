package clients

import (
	"github.com/oliver7100/user-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceClientConfig struct {
	Url string
}

func NewUserServiceClient(config *UserServiceClientConfig) (proto.UserServiceClient, error) {
	c, err := grpc.Dial(config.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return proto.NewUserServiceClient(c), nil
}
