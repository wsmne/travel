package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var Db *gorm.DB

func InitDataBase() error {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/travel?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:@tcp(127.0.0.1:3306)/travel?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("open db error ,err = %v" + err.Error())
		panic(any("db creat error"))
	}
	Db = db
	sqlDB, _ := Db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(50)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(50)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	log.Println("open db success")

	log.Println("数据库迁移开始")
	Db.AutoMigrate(&User{})
	Db.AutoMigrate(&Scene{})
	Db.AutoMigrate(&Score{})
	log.Println("数据库迁移完成")

	scenes := []Scene{
		// 广东景点
		{
			Name:     "广州塔",
			City:     "广州",
			Province: "广东",
			Price:    "150",
			Image:    "https://www.cantontower.com/movies/picture/",
		},
		{
			Name:     "丹霞山",
			City:     "韶关",
			Province: "广东",
			Price:    "120",
			Image:    "https://www.vcg.com/creative-image/danxiashan/",
		},
		{
			Name:     "长隆野生动物世界",
			City:     "广州",
			Province: "广东",
			Price:    "280",
			Image:    "https://www.chimelong.com/gz/safaripark/",
		},

		// 浙江景点
		{
			Name:     "西湖",
			City:     "杭州",
			Province: "浙江",
			Price:    "免费",
			Image:    "https://www.vcg.com/creative-image/xihu/",
		},
		{
			Name:     "乌镇",
			City:     "嘉兴",
			Province: "浙江",
			Price:    "150",
			Image:    "https://www.vcg.com/creative-image/wuzhen/",
		},
		{
			Name:     "千岛湖",
			City:     "杭州",
			Province: "浙江",
			Price:    "180",
			Image:    "https://www.vcg.com/creative-image/qiandaohu/",
		},

		// 江苏景点
		{
			Name:     "中山陵",
			City:     "南京",
			Province: "江苏",
			Price:    "免费",
			Image:    "https://www.vcg.com/creative-image/zhongshanling/",
		},
		{
			Name:     "瘦西湖",
			City:     "扬州",
			Province: "江苏",
			Price:    "100",
			Image:    "https://www.vcg.com/creative-image/shouxihu/",
		},
		{
			Name:     "周庄",
			City:     "苏州",
			Province: "江苏",
			Price:    "100",
			Image:    "https://www.vcg.com/creative-image/zhouzhuang/",
		},
		{
			Name:     "拙政园",
			City:     "苏州",
			Province: "江苏",
			Price:    "90",
			Image:    "https://www.vcg.com/creative-photo/zhuozhengyuan/",
		},
	}
	for i := 0; i < len(scenes); i++ {
		Db.Debug().Create(&scenes[i])
	}
	return nil
}
