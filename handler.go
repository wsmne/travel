package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"proj/travel/controllers"
	"proj/travel/middleware"
)

func RegisterRouter(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		},
		)
	})
	r.POST("/regist", controllers.Register)
	r.GET("/login", controllers.Login)

	auth := r.Group("/api", middleware.JWTMiddleware())
	{
		auth.PUT("/user", controllers.UpdateUser)
	}

}
