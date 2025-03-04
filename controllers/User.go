package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"proj/travel/models"
)

func Register(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 15003,
			"msg":  "输入错误",
		})
		return
	}
	err = models.CreateUser(user)
	if err != nil {
		log.Println("创建用户错误，err : " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 15002,
			"msg":  "创建失败",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  "创建成功",
	})
	return
}

func Login(ctx *gin.Context) {

	return
}

func UpdateUser(ctx *gin.Context) {
	return
}
