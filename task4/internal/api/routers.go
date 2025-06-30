package api

import (
	"github.com/gin-gonic/gin"
	"task4/internal/service"
)

// /api/cms/hellow
const (
	rootpath   = "/api/"
	noAuthPath = "/out/api/"
)

func CmsRouters(r *gin.Engine) {
	cmsApp := service.NewCmsApp()
	sessionAuth := &SessionAuth{}
	jwtTokenAuth := &JWTTokenAuth{}
	// .Use(sessionAuth.Auth) root组的所有接口都需要经过session_id鉴权
	// Use(jwtTokenAuth.AuthJWTToken)  root组的所有接口都需要经过JWT鉴权
	root := r.Group(rootpath).Use(sessionAuth.Auth).Use(jwtTokenAuth.AuthJWTToken)
	{
		root.GET("/cms/hellow", cmsApp.Hellow)
		// 创建文章
		root.POST("/cms/creatPost", cmsApp.CreatePost)
		// 查询所有文章信息
		root.POST("/cms/getAllPosts", cmsApp.GetAllPosts)
		// 根据Id更新文章信息
		root.POST("/cms/updatePostById", cmsApp.UpdatePostById)
		// 根据Id更新文章信息
		root.POST("/cms/deletePostById", cmsApp.DeletePostById)
		// 新增文章评论
		root.POST("/cms/creatComment", cmsApp.CreatComment)
		// 根据文章Id获取该文章的所有评论
		root.POST("/cms/GetAllCommentByPostId", cmsApp.GetAllCommentByPostId)
	}

	noAuth := r.Group(noAuthPath)
	{
		noAuth.POST("/cms/login", cmsApp.Login)
		// 用户注册
		noAuth.POST("/cms/register", cmsApp.Register)
	}

}
