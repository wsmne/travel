package models

import (
	"errors"
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

func FindSceneByID(sceneID uint) (*Scene, error) {
	var scene Scene
	if err := Db.First(&scene, sceneID).Error; err != nil {
		return nil, errors.New("scene not found")
	}
	return &scene, nil
}

func UpdateScene(scene *Scene) (*Scene, error) {
	if err := Db.Save(scene).Error; err != nil {
		return nil, err
	}
	return scene, nil
}
