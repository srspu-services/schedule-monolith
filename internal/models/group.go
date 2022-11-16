package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	GroupName string
	FullName  string
	Course    int
}
