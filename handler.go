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
	r.POST("/register", controllers.Register)
	r.GET("/login", controllers.Login)

	auth := r.Group("/api", middleware.JWTMiddleware())
	{
		auth.PUT("/user", controllers.UpdateUser)
		auth.GET("/recommend", controllers.Recommend)
		auth.GET("/guessULike", controllers.UserFilterRecommend)
		auth.GET("/mostgoods", controllers.MostGoods)
		auth.GET("/mostviews", controllers.MostViews)
		auth.POST("/score", controllers.AddScore)
		auth.POST("/scene/cnt", controllers.Add1)
	}

}
