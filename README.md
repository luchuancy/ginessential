# Gin+Vue前后端分离项目

## 1. 配置环境

### 1.1 go mod 初始化

```go
go mod init cyul.stu0323/ginessential
go mod tidy
```

### 1.2 Gin依赖

```go
go get -u github.com/gin-gonic/gin
```

```go
import "github.com/gin-gonic/gin"
```

### 1.3 测试Gin环境

```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
```



## 2. 实现用户注册

### 2.1 获取用户参数

```go
name := c.PostForm("name")
telephone := c.PostForm("telephone")
password := c.PostForm("password")
```

### 2.2 数据验证

```go
//  用户名为空，生成10位随机字符串
if len(name) == 0 {
    name = util.RandomString(10)
}
// 验证手机号, 返回JSON
if len(telephone) != 11 {
    c.JSON(http.StatusUnprocessableEntity, gin.H{
        "code": 422,
        "msg":  "手机号必须为11位",
    })
    return
}
//  验证密码, 返回JSON
// http.StatusUnprocessableEntity 422状态码
if len(password) <= 6 {
    c.JSON(http.StatusUnprocessableEntity, gin.H{
        "code": 422,
        "msg":  "密码不能少于6位",
    })
    return
}
```

## 3. 连接数据库

### 3.1 下载Gorm库, mysql驱动 

```
go get -u github.com/jinzhu/gorm
go get github.com/go-sql-driver/mysql
```

### 3.2连接数据库

## 4. 实现用户登录

### 4.1 登录路由



### 4.2 Jwt配合中间间用户认证

```go
go get -u github.com/bwplotka/go-jwt
```



##  git

