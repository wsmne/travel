package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var Db *gorm.DB

func InitDataBase() error {
	dsn := "root:@tcp(127.0.0.1:3306)/travel?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("open db error ,err = %v" + err.Error())
		panic("db creat error")
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

	return nil
}
