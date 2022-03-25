package controller

// 用户参数包
import (
	"log"
	"net/http"

	"cyul.stu0323/ginessential/common"
	"cyul.stu0323/ginessential/dto"
	"cyul.stu0323/ginessential/model"
	"cyul.stu0323/ginessential/response"
	"cyul.stu0323/ginessential/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	DB := common.GetDB()
	// 1. 获取参数
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	// 2. 数据验证
	// 2.1 用户名为空，生成10位随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	// 2.2 验证手机号, 返回JSON
	if len(telephone) != 11 {
		// c.JSON(http.StatusUnprocessableEntity, gin.H{
		// 	"code": 422,
		// 	"msg":  "手机号必须为11位",
		// })
		// 格式化返回信息
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	// 2.3 验证密码, 返回JSON
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	// 打印参数
	log.Println(name, telephone, password)
	// 3. 创建用户, 连接数据库 "cyul.stu0323/ginessential/common"
	// 3.1判断手机号是否存在
	if isTelephoneExists(DB, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}
	// 密码加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	// 3.2 不存在则创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)
	// 4. 返回结果 
	response.Success(c, nil, "注册成功")
}

// 登录
func Login(c *gin.Context) {
	DB := common.GetDB()
	// 1. 获取参数
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	// 2. 数据验证
	if len(telephone) != 11 { 
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	// 3. 手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 { 
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	// 4. 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "密码错误",
		})
		return
	}
	// 5. 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		log.Printf("token generate error : %v", err)
		return
	}
	// 6. 返回结果 
	response.Success(c, gin.H{"token": token}, "登录成功")
}

// 获取用户信息
func Info(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"user": dto.ToUserDto(user.(model.User))}, // 类型转换user
	})
}

// 用户手机号是否存在
func isTelephoneExists(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
