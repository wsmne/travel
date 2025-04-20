package models

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var Rdb *redis.Client
var Ctx = context.Background()

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 地址
		Password: "",               // 无密码则留空
		DB:       0,                // 默认使用 0 号数据库
	})

	// 测试连接
	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		panic(any("无法连接 Redis: " + err.Error()))
	}
}

func StartRatingMatrixScheduler() {
	go func() {
		for {
			// 计算到下一个 0 点的时间
			now := time.Now()
			next := now.Add(24 * time.Hour).Truncate(24 * time.Hour)
			duration := next.Sub(now)

			log.Printf("[Scheduler] 已更新矩阵，下次更新在 %s 后更新评分矩阵\n", duration)

			CacheRatingMatrixToRedis()

			time.Sleep(duration) // 睡到明天 0 点
			CacheRatingMatrixToRedis()
			// 接下来每 24 小时执行一次（保险起见再加 ticker）
			ticker := time.NewTicker(24 * time.Hour)
			for range ticker.C {
				CacheRatingMatrixToRedis()
			}
		}
	}()
}
