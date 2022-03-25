package middleware

import (
	"net/http"
	"strings"

	"cyul.stu0323/ginessential/common"
	"cyul.stu0323/ginessential/model"
	"github.com/gin-gonic/gin"
)

// 中间件保护用户信息
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.GetHeader("Authorization")

		// validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg": "权限不足",
			})
			// 抛弃请求
			c.Abort()
			return 
		}
		tokenString = tokenString[7: ]
		// 解析token 
		token, claims, err := common.ParseToken(tokenString) 
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg": "权限不足",
			})
			c.Abort()
			return 
		}

		// 验证通过，获取token的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 验证用户是否存在
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg": "权限不足",
			})
			c.Abort()
			return
		}

		// 用户存在，用户信息写入上下文
		c.Set("user", user)

		c.Next()
	}
}