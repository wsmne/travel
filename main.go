package main

import (
	"github.com/gin-gonic/gin"
	"proj/travel/models"
)

func main() {

	models.InitDataBase()

	r := gin.Default()

	RegisterRouter(r)

	r.Run(":5001")

}
