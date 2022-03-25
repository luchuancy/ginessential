package main

import (
	"cyul.stu0323/ginessential/common"
	"cyul.stu0323/ginessential/router"
	"github.com/gin-gonic/gin"
)

func main() {
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	router.ControlRouter(r)
	panic(r.Run())
}
