package util

import "pro01/grpc/user/proto/service"

var UserServiceClient service.UserServiceClient

func InitCli() {
	UserServiceClient = getUserServiceCli(":9091")
}
