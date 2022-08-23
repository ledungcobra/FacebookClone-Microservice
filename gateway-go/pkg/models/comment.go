package models

import "gorm.io/gorm"

type Comment struct {
	PostID    uint
	Comment   string `gorm:"column:comment"`
	Image     string `gorm:"column:image"`
	CommentBy uint
	Author    User `gorm:"foreignKey:CommentBy"`
	gorm.Model
}
