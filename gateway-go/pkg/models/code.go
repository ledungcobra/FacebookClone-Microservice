package models

import "gorm.io/gorm"

type Code struct {
	Code   string `gorm:"column:code;not null;type:varchar(300);default:''"`
	UserID uint   `gorm:"column:user_id;not null;"`
	gorm.Model
}
