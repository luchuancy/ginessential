package router

import (
	"cyul.stu0323/ginessential/controller"
	"cyul.stu0323/ginessential/middleware"
	"github.com/gin-gonic/gin"
)

func ControlRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)

	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
