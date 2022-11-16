package model

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	GroupID   uint
	UpWeek    bool
	WeekDay   int
	TimeStart time.Time
	TimeEnd   time.Time
	Subject   string
	Classroom string
	Teacher   string
	Type      SubjectType
	Group     Group `gorm:"foreignKey:GroupID"`
}

type InputSchedule struct {
	Week       int    `json:"week"`
	Day        int    `json:"day"`
	Class      int    `json:"class"`
	Auditorium string `json:"auditorium"`
	Lecturer   string `json:"lecturer"`
	Discipline string `json:"discipline"`
	Type       string `json:"type"`
}
