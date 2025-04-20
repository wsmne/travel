package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"proj/travel/models"
	"strings"
	"time"
)

func Add1(c *gin.Context, sceneId uint, addtype, cnt int) {
	scene, err := models.FindSceneByID(sceneId)
	if err != nil {
		return
	}

	if addtype == 1 {
		scene.Goods += cnt
	} else {
		scene.Views += cnt
		userid, _ := c.Get("userid")
		userId := cast.ToUint(userid)
		models.AddUserLog(&models.ClickLog{
			UserId:    userId,
			SceneId:   scene.ID,
			CreatedAt: time.Now(),
		})
	}
	if _, err := models.UpdateScene(scene); err != nil {
		return
	}

}

func GetScenesByType(c *gin.Context) {
	sceneType := c.Query("type")
	city := c.Query("city")
	name := c.Query("name")

	var scenes []models.Scene
	var err error
	switch sceneType {
	case "hot":
		scenes, err = MostGoods(c)
	case "likes":
		scenes, err = MostGoods(c)
	case "views":
		scenes, err = MostViews(c)
	case "guess":
		scenes, err = UserFilterRecommend(c)
	case "recommend":
		scenes, err = MostGoods(c)
	default:
		c.JSON(400, gin.H{"msg": "invalid type"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
	}

	// 对结果进行城市或名称筛选（模糊匹配，支持复合）
	filtered := make([]models.Scene, 0)
	for _, s := range scenes {
		matchCity := city == "" || containsIgnoreCase(s.City, city)
		matchName := name == "" || containsIgnoreCase(s.Name, name)
		if matchCity && matchName {
			filtered = append(filtered, s)
		}
	}

	c.JSON(200, filtered)
}

func containsIgnoreCase(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}

func GetSceneByID(ctx *gin.Context) {
	sceneid := ctx.Param("id")
	sceneId := cast.ToUint(sceneid)
	userid, ok := ctx.Get("userid")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "用户未认证",
		})
		return
	}
	Add1(ctx, sceneId, 2, 1)
	userId := cast.ToUint(userid)
	detail, err := models.FindSceneByID(sceneId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "景点未找到",
		})
		return
	}
	score := models.GetScoreByUserAndSceneId(userId, sceneId)

	ctx.JSON(http.StatusOK, gin.H{
		"code":  0,
		"scene": detail,
		"score": score,
	})
}
