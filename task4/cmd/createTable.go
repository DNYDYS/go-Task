package main

import (
	"fmt"
	"gorm.io/gorm"
	util "task4/internal/utils"
)

type User struct {
	gorm.Model
	Username         string `gorm:"unique;not null"`
	Password         string `gorm:"not null"`
	IsAuthentication string `gorm:"not null"`
	Email            string `gorm:"unique;not null"`
}

type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint
	User    User
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	User    User
	PostID  uint
	Post    Post
}

func main() {
	//db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	db := util.ConnDB()
	// 自动迁移模型
	err := db.Debug().AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		fmt.Println(err)
		return
	}
}
