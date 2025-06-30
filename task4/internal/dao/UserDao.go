package dao

import (
	"fmt"
	"gorm.io/gorm"
	"task4/internal/model"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (userDao *UserDao) GetUserInfoByUserid(userId uint) (*model.User, error) {
	var user model.User
	if err := userDao.db.Debug().Where("id = ?", userId).First(&user).Error; err != nil {
		fmt.Println("GetUserInfoByUserid err:", err)
		return nil, err
	}
	return &user, nil
}

func (userDao *UserDao) GetUserInfoByUserName(userName string) (*model.User, error) {
	var user model.User
	if err := userDao.db.Debug().Where("username = ?", userName).First(&user).Error; err != nil {
		fmt.Println("GetUserInfoByUserName err:", err)
		return nil, err
	}
	return &user, nil
}

// 校验用户是否存在
func (userDao *UserDao) IsExit(userName string) (bool, error) {
	var user model.User

	err := userDao.db.Debug().Where("username", &userName).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		fmt.Println("UserDao IsExit = ", err)
		return false, err
	}

	return true, nil
}

// 插入新的数据
func (userDao *UserDao) Create(user model.User) error {
	if err := userDao.db.Debug().Create(&user).Error; err != nil {
		fmt.Println("UserDao Create user = ", user)
		fmt.Println("UserDao Create = ", err)
		return err
	}
	return nil
}
