package models

import (
	"gorm.io/gorm"
)

type Scene struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null" json:"name"`
	City     string `gorm:"type:varchar(20);not null" json:"city"`
	Province string `gorm:"type:varchar(20);not null" json:"province"`
	Price    string `gorm:"type:varchar(20)"  json:"price"`
	Image    string `gorm:"type:text;not null" json:"image"`
	Goods    int    `gorm:"type:int;not null" json:"goods"`
	Views    int    `gorm:"type:int;not null" json:"views"`
}
