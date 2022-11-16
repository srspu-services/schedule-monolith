package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	MediaId string
	GroupID uint
	Group   Group `gorm:"foreignKey:GroupID"`
}
