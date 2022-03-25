package router

import (
	"cyul.stu0323/ginessential/controller"
	"github.com/gin-gonic/gin"
)

func ControlRouter(r *gin.Engine) {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
}
