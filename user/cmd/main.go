package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"pro01/grpc/user/config"
	"pro01/grpc/user/dao"
	"pro01/grpc/user/proto/service"
)

func main() {
	config.InitConfig()
	err := dao.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	service.RegisterUserServiceServer(server, &service.UserService{})
	listener, err := net.Listen("tcp", "localhost:9091")
	if err != nil {
		log.Fatal(err)
	}
	server.Serve(listener)
}
