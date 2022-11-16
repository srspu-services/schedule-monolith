package model

import "gorm.io/gorm"

type Faculty struct {
	gorm.Model
	FacultyId string
	Code      string
	Name      string
}

type InputFaculty struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
