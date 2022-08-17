package models

import "gorm.io/gorm"

type Post struct {
	AuthorID uint 
	gorm.Model
}
