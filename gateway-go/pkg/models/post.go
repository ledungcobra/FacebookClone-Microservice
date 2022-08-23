package models

import "gorm.io/gorm"

type Post struct {
	Author     User `gorm:"foreignKey:AuthorID"`
	AuthorID   uint
	Type       string    `gorm:"column:type"`
	Text       string    `gorm:"column:text"`
	Images     []string  `gorm:"column:image;serializer:json;type:varchar(500)"`
	Background string    `gorm:"column:background;type:varchar(255)"`
	Comments   []Comment `gorm:"foreignKey:PostID"`
	gorm.Model
}
