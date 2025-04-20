package main

import (
	"github.com/gin-gonic/gin"
	"proj/travel/models"
)

func main() {

	models.InitDataBase()
	models.InitRedis()
	models.StartRatingMatrixScheduler()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*") // 加载 templates 目录下的 HTML 页面
	r.Static("/static", "./static")
	RegisterRouter(r)

	r.Run(":5001")

}
