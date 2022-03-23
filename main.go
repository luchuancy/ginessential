package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gin-gonic/gin"

)

type User struct {
	gorm.Model
	Name string `grom: "type: varchar(20);not null"`
	Telephone string `grom: "varchar(100); not null; unique"`
	Password string `grom: "size: 255; not null"`
}

func main() {
	db := InitDB()
	defer db.Close()

	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		// 获取参数
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")
		// 数据验证
		if len(telephone) != 11 { // 422 or http 常量; gin.H map[string]interface{}
			ctx.JSON(http.StatusUnprocessableEntity, gin.H {
				"code": 422,
				"msg":  "手机号必须为11位",
			}) 
			return
		}
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H {
				"code": 422,
				"msg":  "密码不能少于6位",
			}) 
			return
		}
		// 如果名称没有传，给一个10位字符串
		if len(name) == 0 {
			name = RandomString(10) 
		}

		log.Println(name, telephone, password)
		
		// 判断手机号是否存在
		if isTelephoneExists(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H {
				"code": 422,
				"msg":  "用户已经存在",
			}) 
			return
		}
		// 创建用户
		newUser := User {
			Name: name,
			Telephone: telephone,
			Password: password,
		}
		db.Create(&newUser)
		// 返回结果
		ctx.JSON(200, gin.H {
			"msg": "注册成功",
		})
	})
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}

func isTelephoneExists(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

// 生成随机字符串
func RandomString(n int) string {
	var letters = []byte("asdfghjklqwertyuiopzxcvbnmASDFGHJKLQWERTYUIOPZXCVBNM")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := ""
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", 
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connected database, err: " + err.Error())
	}

	// 自动创建数据表
	db.AutoMigrate(&User{})
	return db
}