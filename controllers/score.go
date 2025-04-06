package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"proj/travel/models"
	"time"
)

// 增加评分的处理函数
func AddScore(c *gin.Context) {
	var rating models.Score
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 验证评分合法性
	if rating.Score < 1 || rating.Score > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Score must be between 1 and 5"})
		return
	}

	// 设置评分时间
	rating.ScoreTime = time.Now()

	// 插入评分记录
	if !models.AddScore(rating) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add rating"})
		return
	}
	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "Rating added successfully"})
}
