package router

import (
	"github.com/gin-gonic/gin"
	"pro01/grpc/api-gateway/handler/user"
)

func NewRouter() *gin.Engine {
	engine := gin.Default()
	userGroup := engine.Group("/user")
	{
		userGroup.POST("/register", user.UserRegisterHandler)
		userGroup.POST("/login", user.UserLoginHandler)
		userGroup.POST("/changeInfo", user.ChangeUserInfoHandler)
		userGroup.GET("/getInfo", user.GetUserInfoHandler)
		userGroup.POST("/delete", user.DeleteUser)
	}
	return engine
}
