package util

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"pro01/grpc/user/proto/service"
)

func getUserServiceCli(addr string) service.UserServiceClient {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	//defer conn.Close()
	client := service.NewUserServiceClient(conn)
	return client
}
