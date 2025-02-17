package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	UserName string `gorm:"type:varchar(20);not null"`
	UserPW   string `gorm:"type:varchar(20);not null"`
}
