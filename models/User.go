package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UID      string `gorm:"primaryKey column:uid" json:"uid"`
	UserName string `gorm:"type:varchar(20);not null" json:"username"`
	UserPW   string `gorm:"type:varchar(20);not null" json:"password"`
}

func GetUserByUID(uid string) (user User, err error) {
	err = Db.Debug().Where("uid = ?", uid).First(&user).Error
	return user, err
}

func CreateUser(user User) error {
	err := Db.Debug().Create(&user).Error
	return err
}

func GetUserByName(name string) (user User, err error) {
	err = Db.Debug().Where("user_name = ?", name).First(&user).Error
	return user, err
}
