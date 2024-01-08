package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Uid       uint   `gorm:"not null" json:"uid"`
	Title     string `json:"title"`
	Status    int    `gorm:"default:0" json:"status"`
	Content   string `gorm:"type:longtext" json:"content"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}
