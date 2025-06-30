package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint
	User    User
	Comment []Comment
}

func (p Post) TableName() string {
	table := "cms.posts"
	return table
}
