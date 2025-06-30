package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"task4/internal/dao"
	"task4/internal/model"
)

type RegisterReq struct {
	UserId           string `json:"userId" `
	Username         string `json:"username" binding:"required"`
	Password         string `json:"password" binding:"required"`
	Email            string `json:"email" binding:"required"`
	IsAuthentication string `json:"isAuthentication" binding:"required"`
}

type RegisterResp struct {
	Message string `json:"message" binding:"required"`
}

func (c *CmsApp) Register(ctx *gin.Context) {
	var user RegisterReq
	userInfo := &user
	if err := ctx.ShouldBindJSON(userInfo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("注册接口入参 registerReq:", userInfo)

	// 账号校验
	userDao := dao.NewUserDao(c.db)
	isExit, err := userDao.IsExit(userInfo.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if isExit == true {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "用户已存在"})
		return
	}

	// 密码加密
	hashedPwd, err := encrypPwd(userInfo.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("userInfo:", userInfo)
	// 用户信息持久化
	if err := userDao.Create(model.User{
		Username:         userInfo.Username,
		Password:         hashedPwd,
		Email:            userInfo.Email,
		IsAuthentication: userInfo.IsAuthentication,
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data": &RegisterResp{
			Message: fmt.Sprintf("注册成功，你好", userInfo.Username),
		},
	})
}

/*
*
密码加密
*/
func encrypPwd(password string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("密码加密异常：", err)
		return "", err
	}
	return string(hashedPwd), nil
}
