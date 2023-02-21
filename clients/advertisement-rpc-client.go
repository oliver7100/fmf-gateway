package clients

import (
	"github.com/oliver7100/advertisement-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAdvertisementClient(cfg *clientConfig) (proto.AdvertisementServiceClient, error) {
	d, err := grpc.Dial(cfg.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return proto.NewAdvertisementServiceClient(d), nil
}
