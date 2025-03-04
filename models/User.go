package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20);not null" json:"username"`
	UserPW   string `gorm:"type:varchar(20);not null" json:"password"`
}

func GetUserByID(id uint) (user User, err error) {
	Db.AutoMigrate(&user)
	err = Db.Debug().First(&user, id).Error
	return user, err
}
func CreateUser(user User) error {
	Db.AutoMigrate(&user)
	err := Db.Debug().Create(&user).Error
	return err
}
