package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"log"
	"net/http"
	"proj/travel/middleware"
	"proj/travel/models"
)

func Register(ctx *gin.Context) {
	//body, err := io.ReadAll(ctx.Request.Body)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{
	//		"code": 15006,
	//		"msg":  "读取请求体失败",
	//	})
	//	return
	//}
	//fmt.Println(string(body)) // 打印请求体，查看是否是有效的 JSON 数据
	//
	//ctx.Request.Body = io.NopCloser(bytes.NewReader(body))

	var user models.User
	err := ctx.ShouldBindJSON(&user)
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
			"msg":  "用户名已存在",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "创建成功",
	})
	return
}

func Login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 15003,
			"msg":  "输入错误",
		})
		return
	}
	login, err := models.GetUserByUserName(user.UserName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 15004,
			"msg":  "用户不存在",
		})
		return
	}
	if user.UserPW != login.UserPW {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 15005,
			"msg":  "密码错误",
		})
		return
	}
	token, _ := middleware.GenerateJWT(login.ID)
	ctx.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "登录成功",
		"token": token,
	})
	return
}

func UpdateUser(ctx *gin.Context) {
	uid, _ := ctx.Get("uid")
	log.Println(uid)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
	return
}

func GetUserInfo(c *gin.Context) {
	userid, ok := c.Get("userid")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "用户未认证",
		})
		return
	}

	user, err := models.GetUserById(cast.ToUint(userid))
	if user == nil || err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "用户不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_name": user.UserName,
	})
}
