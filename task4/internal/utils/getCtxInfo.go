package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取当前登录用户ID
func GetUserIdFromCtx(ctx *gin.Context) uint {

	id, isExistId := ctx.Get("id")
	fmt.Println(id)
	if !isExistId {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "用户登陆失效，请重新登录"})
		return 0
	}

	userId, err := GetUintID(ctx, "id")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}
	fmt.Println("GetUintID加载userId的值：", userId)
	return userId
}
