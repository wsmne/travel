package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"math"
	"net/http"
	"proj/travel/models"
	"sort"
)

func MostViews(c *gin.Context) {
	result, err := models.GetTopKScenesByViews(10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load scenes from db"})
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func MostGoods(c *gin.Context) {
	result, err := models.GetTopKScenesByGoods(10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load scenes from db"})
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func UserFilterRecommend(ctx *gin.Context) {
	userId := ctx.Param("userid")
	result := recommendScenesByUser(cast.ToUint(userId), 2)
	ctx.JSON(200, result)
}
func Recommend(ctx *gin.Context) {

}

func buildRatingMatrix() map[uint]map[uint]float64 {
	var scores []models.Score
	models.Db.Find(&scores)

	ratingMatrix := make(map[uint]map[uint]float64)
	for _, s := range scores {
		if _, ok := ratingMatrix[s.UserId]; !ok {
			ratingMatrix[s.UserId] = make(map[uint]float64)
		}
		ratingMatrix[s.UserId][s.SceneId] = float64(s.Score)
	}
	return ratingMatrix
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
	ratings := buildRatingMatrix()
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
	recommendCount := 5
	if len(predictions) < recommendCount {
		recommendCount = len(predictions)
	}

	var sceneIds []uint
	for i := 0; i < recommendCount; i++ {
		sceneIds = append(sceneIds, predictions[i].SceneId)
	}

	var scenes []models.Scene
	models.Db.Where("id IN ?", sceneIds).Find(&scenes)
	return scenes
}
