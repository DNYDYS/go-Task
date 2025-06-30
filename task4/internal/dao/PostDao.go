package dao

import (
	"fmt"
	"gorm.io/gorm"
	"task4/internal/model"
	util "task4/internal/utils"
)

type PostDao struct {
	db *gorm.DB
}

func NewPostDao(db *gorm.DB) *PostDao {
	return &PostDao{db: db}
}

/*
*新增一篇文章
 */
func (postDao *PostDao) CreatPost(post model.Post) error {
	if err := postDao.db.Debug().Create(&post).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (postDao *PostDao) GetAllPosts(post model.Post) *gorm.DB {
	var posts []model.Post
	query := postDao.db.Debug().Model(&posts)
	if post.ID != 0 {
		query = query.Where("posts.id = ?", post.ID)
	}
	if !util.IsNilOrEmptyString(&post.Title) {
		query = query.Where("posts.title like ?", "%"+post.Title+"%")
	}
	if !util.IsNilOrEmptyString(&post.Content) {
		query = query.Where("posts.content like ?", "%"+post.Content+"%")
	}

	return query
}

func (postDao *PostDao) GetPostInfoById(id uint) (*model.Post, error) {
	var post model.Post
	if err := postDao.db.Debug().Where("id = ?", id).Take(&post).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &post, nil

}

func (postDao *PostDao) UpdatePostById(post model.Post) error {
	if err := postDao.db.Debug().Updates(&post).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (postDao *PostDao) DeletePostById(id uint) error {
	if err := postDao.db.Debug().Where("id = ?", id).Delete(&model.Post{}).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
