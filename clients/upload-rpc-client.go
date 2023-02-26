package clients

import (
	"context"
	"io"
	"mime/multipart"
	"time"

	"github.com/oliver7100/upload-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUploadServiceClient(cfg *clientConfig) (proto.UploadServiceClient, error) {
	c, err := grpc.Dial(cfg.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return proto.NewUploadServiceClient(c), nil
}

func UploadFile(uploadServiceClient proto.UploadServiceClient, file *multipart.FileHeader) (*string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := uploadServiceClient.UploadImage(ctx)
	if err != nil {
		return nil, err
	}

	req := proto.UploadImageRequest{
		Data: &proto.UploadImageRequest_Info{
			Info: &proto.ImageInfo{
				Type: "jpeg",
			},
		},
	}

	err = stream.Send(&req)
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, 1024)

	oFile, _ := file.Open()

	for {
		num, err := oFile.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		if err := stream.Send(&proto.UploadImageRequest{Data: &proto.UploadImageRequest_ChunkData{ChunkData: buffer[:num]}}); err != nil {
			break
		}
	}

	res, _ := stream.CloseAndRecv()

	return &res.Uri, nil
}
