package models

import (
	"gorm.io/gorm"
)

type Score struct {
	gorm.Model
	UserId     uint    `gorm:"not null" json:"user_id"`
	SceneId    uint    `gorm:"not null" json:"scene_id"`
	Score      float64 `gorm:"not null" json:"score"`
	ScoreType  int     `gorm:"not null" json:"score_type"`
	ScoreTime  string  `gorm:"not null" json:"score_time"`
	ScoreDesc  string  `json:"score_desc"`
	ScoreState int     `json:"score_state"`
}
