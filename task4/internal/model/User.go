package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username         string `gorm:"unique;not null"`
	Password         string `gorm:"not null"`
	IsAuthentication string `gorm:"not null"`
	Email            string `gorm:"unique;not null"`
}

func (u User) TableName() string {
	table := "cms.users"
	return table
}
