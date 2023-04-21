package service

import (
	"context"
	"pro01/grpc/user/dao"
)

type UserService struct {
}

var userService UserServiceServer = &UserService{}

func (u *UserService) UserRegister(ctx context.Context, request *UserRequest) (*UserResponse, error) {
	err := dao.UserRegister(request.User.UserId, request.User.Pwd)
	if err != nil {
		return nil, err
	}
	return &UserResponse{Code: 0, Msg: "success"}, nil
}

func (u *UserService) UserLogin(ctx context.Context, request *UserRequest) (*UserResponse, error) {
	err := dao.UserLogin(request.User.UserId, request.User.Pwd)
	if err != nil {
		return nil, err
	}
	return &UserResponse{Code: 0, Msg: "success"}, nil
}

func (u *UserService) UserChangeInfo(ctx context.Context, request *UserRequest) (*UserResponse, error) {
	err := dao.ChangeUserInfo(request.User.UserId, request.User.Nickname, request.User.Signature)
	if err != nil {
		return nil, err
	}
	return &UserResponse{Code: 0, Msg: "success"}, nil
}

func (u *UserService) UserGetInfo(ctx context.Context, request *UserRequest) (*UserResponse, error) {
	user, err := dao.GetUserInfo(request.User.UserId)
	if err != nil {
		return nil, err
	}
	return &UserResponse{Code: 0, Msg: "success", User: &UserModel{
		UserId:    user.UserId,
		Pwd:       user.Pwd,
		Nickname:  user.Nickname,
		Signature: user.Signature},
	}, nil
}

func (u *UserService) UserDelete(ctx context.Context, request *UserRequest) (*UserResponse, error) {
	err := dao.DeleteUser(request.User.UserId)
	if err != nil {
		return nil, err
	}
	return &UserResponse{Code: 0, Msg: "success"}, nil
}

func (u *UserService) mustEmbedUnimplementedUserServiceServer() {

}
