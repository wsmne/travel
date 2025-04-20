package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"math"
	"proj/travel/models"
	"sort"
)

func MostViews(c *gin.Context) ([]models.Scene, error) {
	result, err := models.GetTopKScenesByViews(10)
	if err != nil {
		return nil, fmt.Errorf("加载失败")
	}
	return result, nil
}

func MostGoods(c *gin.Context) ([]models.Scene, error) {
	result, err := models.GetTopKScenesByGoods(100)
	if err != nil {
		return nil, fmt.Errorf("加载失败")
	}
	return result, nil
}

func UserFilterRecommend(ctx *gin.Context) ([]models.Scene, error) {
	userId, ok := ctx.Get("userid")
	if !ok {
		return nil, nil
	}
	result := recommendScenesByUser(cast.ToUint(userId), 2)
	return result, nil
}
func Recommend(ctx *gin.Context) {

}

func pearsonSimilarity(a, b map[uint]float64) float64 {
	var common []uint
	for item := range a {
		if _, ok := b[item]; ok {
			common = append(common, item)
		}
	}

	n := float64(len(common))
	if len(common) == 0 {
		return 0
	}

	var sumA, sumB, sumA2, sumB2, sumAB float64
	for _, item := range common {
		rA := a[item]
		rB := b[item]

		sumA += rA
		sumB += rB
		sumA2 += rA * rA
		sumB2 += rB * rB
		sumAB += rA * rB
	}

	numerator := sumAB - (sumA * sumB / n)
	denominator := math.Sqrt((sumA2 - (sumA * sumA / n)) * (sumB2 - (sumB * sumB / n)))

	if denominator == 0 {
		return 0
	}

	return numerator / denominator
}

type SimilarUser struct {
	UserId     uint
	Similarity float64
}

func findTopKSimilarUsers(targetUser uint, ratings map[uint]map[uint]float64, k int) []SimilarUser {
	var similarUsers []SimilarUser

	for userId, rating := range ratings {
		if userId == targetUser {
			continue
		}
		sim := pearsonSimilarity(ratings[targetUser], rating)
		if sim > 0 {
			similarUsers = append(similarUsers, SimilarUser{UserId: userId, Similarity: sim})
		}
	}

	sort.Slice(similarUsers, func(i, j int) bool {
		return similarUsers[i].Similarity > similarUsers[j].Similarity
	})

	if len(similarUsers) > k {
		return similarUsers[:k]
	}
	return similarUsers
}

func predictScores(targetUser uint, ratings map[uint]map[uint]float64, similarUsers []SimilarUser) map[uint]float64 {
	predicted := make(map[uint]float64)
	weightSum := make(map[uint]float64)

	for _, neighbor := range similarUsers {
		for sceneId, score := range ratings[neighbor.UserId] {
			if _, rated := ratings[targetUser][sceneId]; rated {
				continue // 跳过已评分
			}

			predicted[sceneId] += neighbor.Similarity * score
			weightSum[sceneId] += math.Abs(neighbor.Similarity)
		}
	}

	for sceneId := range predicted {
		if weightSum[sceneId] > 0 {
			predicted[sceneId] /= weightSum[sceneId]
		}
	}

	return predicted
}

type Prediction struct {
	SceneId uint
	Score   float64
}

func recommendScenesByUser(userId uint, k int) []models.Scene {
	ratings := models.GetRatingMatrixFromRedis()
	ratings[userId] = models.GetUserRealTimeRatingsWithClicks(userId)

	topUsers := findTopKSimilarUsers(userId, ratings, k)
	predicted := predictScores(userId, ratings, topUsers)

	var predictions []Prediction
	for sceneId, score := range predicted {
		predictions = append(predictions, Prediction{SceneId: sceneId, Score: score})
	}

	sort.Slice(predictions, func(i, j int) bool {
		return predictions[i].Score > predictions[j].Score
	})

	// 获取推荐景点详细信息
	recommendCount := 100
	if len(predictions) < recommendCount {
		recommendCount = len(predictions)
	}

	var sceneIds []uint
	for i := 0; i < recommendCount; i++ {
		sceneIds = append(sceneIds, predictions[i].SceneId)
	}

	var scenes []models.Scene
	models.Db.Where("id IN ?", sceneIds).Find(&scenes)

	// 如果推荐的景点少于预期数量，从热门景点中补充
	if len(scenes) < recommendCount {
		// 获取热门景点的 ID
		topScenes, err := models.GetTopKScenesByGoods(recommendCount - len(scenes))
		if err != nil {
			return scenes
		}

		// 将热门景点的 ID 加入到推荐列表中
		for _, scene := range topScenes {
			sceneIds = append(sceneIds, scene.ID)
		}

		// 获取热门景点的详细信息
		var additionalScenes []models.Scene
		models.Db.Where("id IN ?", sceneIds).Find(&additionalScenes)
		scenes = append(scenes, additionalScenes...)
	}

	return scenes
}
