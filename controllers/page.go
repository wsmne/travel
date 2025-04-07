package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FirstPage(c *gin.Context) {
	c.HTML(http.StatusOK, "firstpage.html", nil)
}

func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "homepage.html", nil)
}
