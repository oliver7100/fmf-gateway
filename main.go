package main

import (
	"context"
	"fmt"
	"log"

	"github.com/oliver7100/fmf-gateway/clients"
	"github.com/oliver7100/user-service/proto"
)

func main() {
	log.Println("test")

	c, err := clients.NewUserServiceClient(
		&clients.UserServiceClientConfig{
			Url: ":9000",
		},
	)

	if err != nil {
		panic(err)
	}

	res, _ := c.CreateUser(context.Background(), &proto.CreateUserRequest{
		User: &proto.User{
			Name:     "Henning",
			Email:    "Henning@gmail.com",
			Password: "Hennnnnneee",
		},
	})

	fmt.Println(res)
}
