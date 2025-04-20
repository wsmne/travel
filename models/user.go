package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20);not null;unique" json:"username"`
	UserPW   string `gorm:"type:varchar(20);not null" json:"password"`
}

func GetUserByUserName(userName string) (user User, err error) {
	err = Db.Debug().Where("user_name = ?", userName).First(&user).Error
	return user, err
}

func CreateUser(user User) error {
	err := Db.Debug().Create(&user).Error
	return err
}

func GetUserById(id uint) (user *User, err error) {
	err = Db.Debug().Where("id = ?", id).First(&user).Error
	return user, err
}
