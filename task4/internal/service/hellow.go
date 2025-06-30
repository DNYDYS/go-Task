package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *CmsApp) Hellow(ctx *gin.Context) {

	userName, isExit := ctx.Get("username")
	if !isExit {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "用户登陆失效，请重新登录"})
	}
	fmt.Println("Hellow:", userName)

}
