package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"pro01/grpc/api-gateway/util"
	"pro01/grpc/user/proto/service"
)

func UserRegisterHandler(c *gin.Context) {
	userId, pwd := c.PostForm("userId"), c.PostForm("pwd")
	resp, err := util.UserServiceClient.UserRegister(context.Background(), &service.UserRequest{User: &service.UserModel{UserId: userId, Pwd: pwd}})
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code": resp.Code,
			"msg":  resp.Msg,
		})
	}
}

func UserLoginHandler(c *gin.Context) {
	userId, pwd := c.PostForm("userId"), c.PostForm("pwd")
	resp, err := util.UserServiceClient.UserLogin(context.Background(), &service.UserRequest{User: &service.UserModel{UserId: userId, Pwd: pwd}})
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code": resp.Code,
			"msg":  resp.Msg,
		})
	}
}

func ChangeUserInfoHandler(c *gin.Context) {
	userId := c.PostForm("userId")
	nickname, signature := c.PostForm("nickname"), c.PostForm("signature")
	resp, err := util.UserServiceClient.UserChangeInfo(context.Background(), &service.UserRequest{User: &service.UserModel{
		UserId:    userId,
		Nickname:  nickname,
		Signature: signature,
	}})
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code": resp.Code,
			"msg":  resp.Msg,
		})
	}
}

func GetUserInfoHandler(c *gin.Context) {
	userId := c.Query("userId")
	resp, err := util.UserServiceClient.UserGetInfo(context.Background(), &service.UserRequest{User: &service.UserModel{UserId: userId}})
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code": resp.Code,
			"msg":  resp.Msg,
			"user": resp.User,
		})
	}
}

func DeleteUser(c *gin.Context) {
	userId := c.PostForm("userId")
	resp, err := util.UserServiceClient.UserDelete(context.Background(), &service.UserRequest{User: &service.UserModel{UserId: userId}})
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code": resp.Code,
			"msg":  resp.Msg,
		})
	}
}
