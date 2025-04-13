package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FirstPage(ctx *gin.Context) {
	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")
	ctx.HTML(http.StatusOK, "firstpage.html", nil)
}

func HomePage(ctx *gin.Context) {
	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")
	ctx.HTML(http.StatusOK, "homepage.html", nil)
}

func DetailPage(ctx *gin.Context) {
	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")
	ctx.HTML(http.StatusOK, "detail.html", nil)
}
