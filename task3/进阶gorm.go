package main

import (
	"fmt"
	"gorm.io/gorm"
	"task1/util"
)

type User struct {
	Id       uint   `gorm:"primarykey;auto_increment"`
	Username string `gorm:"varchar(20)"`
	Age      string `gorm:"varchar(20)"`
	Posts    []Post `gorm:"foreignkey:UserId"`
}
type Post struct {
	Id         uint      `gorm:"primarykey;auto_increment"`
	Title      string    `gorm:"varchar(200)"`
	Content    string    `gorm:"varchar(2000)"`
	CommentNum uint      `gorm:"int(11)"`
	Status     string    `gorm:"varchar(20)"`
	UserId     uint      `gorm:"int(11)"`
	User       User      `gorm:"foreignkey:UserId"`
	Comment    []Comment `gorm:"foreignkey:PostId"`
}
type Comment struct {
	Id      uint   `gorm:"primarykey;auto_increment"`
	Content string `gorm:"varchar(2000)"`
	PostId  uint   `gorm:"int(11)"`
	UserId  uint   `gorm:"int(11)"`
	Post    Post   `gorm:"foreignkey:PostId"`
	User    User   `gorm:"foreignkey:UserId"`
}

// BeforeCreate 钩子函数
func (c *Comment) BeforeCreate(db *gorm.DB) error {
	return db.Debug().Model(&Post{}).Where("id = ?", c.PostId).Update("comment_num", gorm.Expr("comment_num + ?", 1)).Error
}

// AfterDelete 钩子函数
func (c *Comment) AfterDelete(db *gorm.DB) error {
	var post Post
	if err := db.Debug().Preload("comments").First(&post, c.PostId).Error; err != nil {
		return err
	}

	if err := db.Debug().Model(&Post{}).Where("id = ?", c.PostId).Update("Status", gorm.Expr("Status - ?", 1)).Error; err != nil {
		return err
	}

	if len(post.Comment) == 0 {
		return db.Debug().Model(&Post{}).Where("id = ?", c.PostId).Update("Status", "无评论").Error
	}

	return nil
}

func main() {
	db := util.ConnDB()

	// 题目一
	//使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章），
	//Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
	//编写Go代码，使用Gorm创建这些模型对应的数据库表。
	//db.Debug().AutoMigrate(User{}, Post{}, Comment{})

	createData(db)

	// 题目二
	//编写Go代码，使用Gorm查询某个用户发布的所有文章及其应的评论信息。
	//var posts []Post
	//if err := db.Debug().Model(&posts).Preload("Comment").Where("user_id = ?", "1").Find(&posts).Error; err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(posts)
	//编写Go代码，使用Gorm查询评论数量最多的文章信息。
	var post Post
	if err := db.Debug().Model(post).Order("comment_num DESC").Limit(1).Find(&post).Error; err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(post)

}

func createData(db *gorm.DB) {
	users := []User{
		{
			Username: "nike",
		},
		{
			Username: "jack",
		},
		{
			Username: "marry",
		},
	}
	db.Debug().Create(&users)

	psots := []Post{
		{
			Title:   "java入门到裂开",
			Content: "nike的首个书籍，首秀哦",
			Status:  "0",
			UserId:  1,
		},
		{
			Title:   "go入门到放弃",
			Content: "jack的第二个书籍",
			Status:  "0",
			UserId:  2,
		},
	}

	db.Debug().Create(&psots)

	comments := []Comment{
		{
			Content: "入门到裂开，哈哈哈哈",
			PostId:  5,
			UserId:  3,
		},
		{
			Content: "放弃，怎么可能放弃",
			PostId:  5,
			UserId:  3,
		},
	}
	db.Debug().Create(&comments)

}
