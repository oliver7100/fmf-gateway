package clients

import (
	"context"
	"fmt"
	"io"
	"log"
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

func UploadFile(uploadServiceClient proto.UploadServiceClient, file *multipart.FileHeader) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := uploadServiceClient.UploadImage(ctx)
	if err != nil {
		log.Fatal("cannot upload image: ", err)
	}

	req := proto.UploadImageRequest{
		Data: &proto.UploadImageRequest_Info{
			Info: &proto.ImageInfo{
				Type: "jpg",
			},
		},
	}

	err = stream.Send(&req)
	if err != nil {
		log.Fatal("cannot send image info to server: ", err, stream.RecvMsg(nil))
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

	fmt.Println(res.GetSize())
}
