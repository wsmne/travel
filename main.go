package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	RegisterRouter(r)

	r.Run(":5001")

}
