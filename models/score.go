package models

import (
	"context"
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"math"
	"time"
)

type Score struct {
	gorm.Model
	UserId     uint      `gorm:"not null" json:"user_id"`
	SceneId    uint      `gorm:"not null" json:"scene_id"`
	Score      float64   `gorm:"not null" json:"score"`
	ScoreType  int       `gorm:"not null" json:"score_type"`
	ScoreTime  time.Time `gorm:"not null" json:"score_time"`
	ScoreDesc  string    `json:"score_desc"`
	ScoreState int       `json:"score_state"`
	IsLike     bool      `json:"is_like"`
}

func AddScore(rating Score) bool {
	// 插入评分记录
	if err := Db.Create(&rating).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

func GetScoreByUserAndSceneId(userID uint, scoreID uint) *Score {
	var score Score

	err := Db.Where("user_id = ? AND scene_id = ?", userID, scoreID).First(&score).Error

	if err != nil {
		return nil
	}
	return &score

}

func UpdateScore(score *Score) (*Score, error) {
	if err := Db.Save(score).Error; err != nil {
		return nil, err
	}
	return score, nil
}

func BuildRatingMatrix() map[uint]map[uint]float64 {
	var scores []Score
	Db.Find(&scores)

	ratingMatrix := make(map[uint]map[uint]float64)
	for _, s := range scores {
		if _, ok := ratingMatrix[s.UserId]; !ok {
			ratingMatrix[s.UserId] = make(map[uint]float64)
		}
		ratingMatrix[s.UserId][s.SceneId] = float64(s.Score)
	}
	return ratingMatrix
}

func GetRatingMatrixFromRedis() map[uint]map[uint]float64 {
	val, err := Rdb.Get(context.Background(), "rating_matrix").Result()
	if err != nil {
		log.Println("Redis rating_matrix not found, fallback to DB")
		return BuildRatingMatrix()
	}

	var ratings map[uint]map[uint]float64
	json.Unmarshal([]byte(val), &ratings)
	return ratings
}

func GetUserRealTimeRatingsWithClicks(userId uint) map[uint]float64 {
	var scores []Score
	Db.Where("user_id = ?", userId).Find(&scores)

	ratingMap := make(map[uint]float64)
	for _, s := range scores {
		ratingMap[s.SceneId] = float64(s.Score)
	}

	// 加入点击加权：获取用户近一天点击记录
	var clicks []ClickLog // 你得有个表记录用户点击
	startTime := time.Now().Add(-24 * time.Hour)
	Db.Where("user_id = ? AND created_at >= ?", userId, startTime).Find(&clicks)

	clickWeights := make(map[uint]float64)
	for _, click := range clicks {
		clickWeights[click.SceneId] += 1
	}

	// 点击加权公式
	for sceneId, count := range clickWeights {
		weight := math.Log(float64(count)+1) * 0.3 // 可调的 β 值
		ratingMap[sceneId] += weight
	}

	return ratingMap
}
