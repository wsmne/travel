package main

import (
	"github.com/gin-gonic/gin"
	"proj/travel/controllers"
	"proj/travel/middleware"
)

func RegisterRouter(r *gin.Engine) {
	r.GET("/firstpage", controllers.FirstPage)
	r.GET("/homepage", controllers.HomePage)
	r.GET("/detail", controllers.DetailPage)

	// 注册和登录 POST 请求
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// 保护的路由组，带有 JWT 中间件
	auth := r.Group("/api", middleware.JWTMiddleware())
	{
		auth.GET("/scene", controllers.GetScenesByType)
		auth.GET("/scene/:id", controllers.GetSceneByID)
		auth.POST("/score", controllers.AddScore)
		auth.POST("/score/like", controllers.AddLike)
		//auth.GET("/recommend", controllers.Recommend)
		//auth.GET("/mostviews", controllers.MostViews)
		//auth.GET("/mostgoods", controllers.MostGoods)
		//auth.GET("/guessulike", controllers.UserFilterRecommend)
	}

	//r.POST("/register", controllers.Register)
	//r.GET("/login", controllers.Login)
	//
	//auth := r.Group("/api", middleware.JWTMiddleware())
	//{
	//	auth.PUT("/user", controllers.UpdateUser)
	//	auth.GET("/recommend", controllers.Recommend)
	//	auth.GET("/guessULike", controllers.UserFilterRecommend)
	//	auth.GET("/mostgoods", controllers.MostGoods)
	//	auth.GET("/mostviews", controllers.MostViews)
	//	auth.POST("/score", controllers.AddScore)
	//	auth.POST("/scene/cnt", controllers.Add1)
	//}

}
