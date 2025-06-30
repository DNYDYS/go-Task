package dao

import (
	"fmt"
	"gorm.io/gorm"
	"task4/internal/model"
)

type CommentDao struct {
	db *gorm.DB
}

func NewCommentDao(db *gorm.DB) *CommentDao {
	return &CommentDao{db: db}
}

func (commentDao *CommentDao) CreatComment(comment model.Comment) error {
	if err := commentDao.db.Debug().Create(&comment).Error; err != nil {
		fmt.Println("新增评论异常：", err)
		return err
	}
	return nil
}

func (commentDao *CommentDao) GetAllCommentByPostId(postId uint) ([]model.Comment, error) {
	var comments []model.Comment
	if err := commentDao.db.Debug().Where("post_id = ?", postId).Find(&comments).Error; err != nil {
		fmt.Println("根据文章Id查询对应所有评论信息异常：", err)
		return nil, err
	}

	return comments, nil
}
