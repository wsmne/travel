package models

import (
	"gorm.io/gorm"
	"log"
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
