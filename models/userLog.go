package models

import (
	"context"
	"encoding/json"
	"time"
)

type ClickLog struct {
	ID        uint `gorm:"primaryKey"`
	UserId    uint
	SceneId   uint
	CreatedAt time.Time
}

// 每天凌晨 0 点调用
func CacheRatingMatrixToRedis() {
	ratings := BuildRatingMatrix() // 就是你已有的函数

	data, _ := json.Marshal(ratings)
	Rdb.Set(context.Background(), "rating_matrix", data, 24*time.Hour) // 可设置自动过期
}

func AddUserLog(log *ClickLog) error {
	return Db.Create(log).Error
}
