package models

type Scene struct {
	ID       uint   `gorm:"primaryKey"`
	ScenicID string `gorm:"type:varchar(20);not null"`
	Scenic   string `gorm:"type:varchar(20);not null"`
	City     string `gorm:"type:varchar(20);not null"`
	Province string `gorm:"type:varchar(20);not null"`
	Country  string `gorm:"type:varchar(20);not null"`
	Price    string `gorm:"type:varchar(20)"`
	Image    string `gorm:"type:varchar;not null"`
}
