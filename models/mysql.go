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
	err = Db.AutoMigrate(&User{})
	if err != nil {
		log.Println("auto migrate error,err = %v" + err.Error())
		panic(err)
	}
	err = Db.AutoMigrate(&Scene{})
	if err != nil {
		log.Println("auto migrate error,err = %v" + err.Error())
		panic(err)
	}
	err = Db.AutoMigrate(&Score{})
	if err != nil {
		log.Println("auto migrate error,err = %v" + err.Error())
		panic(err)
	}
	log.Println("数据库迁移完成")

	scenes := []Scene{
		// 广东景点
		{
			Name:     "广州塔",
			City:     "广州",
			Province: "广东",
			Price:    "150",
			Image:    "data:image/jpeg;base64,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",
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
	err = db.Create(&scenes).Error
	if err != nil {
		log.Println("create scene error,err = %v" + err.Error())
		panic(err)
	}
	return nil
}
