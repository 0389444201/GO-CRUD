package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Desc  string
	Body  string
	Title string
	Name  string
}
