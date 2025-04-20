package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"proj/travel/models"
	"time"
)

// 增加评分的处理函数
func AddScore(c *gin.Context) {
	userid, ok := c.Get("userid")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "用户未认证",
		})
		return
	}
	userId := cast.ToUint(userid)
	var rating models.Score
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	rating.UserId = userId

	// 验证评分合法性
	if rating.Score < 1 || rating.Score > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Score must be between 1 and 5"})
		return
	}

	score := models.GetScoreByUserAndSceneId(rating.UserId, rating.SceneId)
	if score == nil {
		rating.ScoreTime = time.Now()
		if !models.AddScore(rating) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add rating"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "Rating added successfully",
		})
		return
	}

	// 设置评分时间
	score.ScoreTime = time.Now()

	models.UpdateScore(score)

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Rating added successfully",
	})
}

// 点赞的处理函数
func AddLike(c *gin.Context) {
	userid, ok := c.Get("userid")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "用户未认证",
		})
		return
	}
	userId := cast.ToUint(userid)
	var rating models.Score
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	rating.UserId = userId

	score := models.GetScoreByUserAndSceneId(rating.UserId, rating.SceneId)
	if score == nil {
		rating.ScoreTime = time.Now()
		if !models.AddScore(rating) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add rating"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "Rating added successfully",
		})
		return
	}
	// 设置评分时间
	score.ScoreTime = time.Now()
	score.IsLike = rating.IsLike
	models.UpdateScore(score)
	if rating.IsLike == true {
		Add1(c, rating.SceneId, 1, 1)
	} else {
		Add1(c, rating.SceneId, 1, -1)
	}
	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "like added successfully",
	})
}
