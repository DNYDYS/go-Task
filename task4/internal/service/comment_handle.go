package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"task4/internal/common"
	"task4/internal/common/logger"
	"task4/internal/dao"
	"task4/internal/model"
	util "task4/internal/utils"
)

type CommentReq struct {
	Content string `json:"content" `
	//UserID  uint   `json:"user_id" binding:"required"`
	PostID uint `json:"post_id" binding:"required"`
}

func (c *CmsApp) CreatComment(ctx *gin.Context) {
	var commentReq CommentReq
	if err := ctx.ShouldBindJSON(&commentReq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("新增文章评论接口入参:", commentReq)
	if util.IsNilOrEmptyString(&commentReq.Content) {
		ctx.Error(common.NewAppError(500, "评论内容不能为空"))
		return
	}

	// 根据id获取文章信息
	postDao := dao.NewPostDao(c.db)
	post, err := postDao.GetPostInfoById(commentReq.PostID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if post == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "文章不存在，无法评论"})
		return
	}

	userId := util.GetUserIdFromCtx(ctx)

	commentDao := dao.NewCommentDao(c.db)
	createErr := commentDao.CreatComment(model.Comment{
		Content: commentReq.Content,
		UserID:  userId,
		PostID:  commentReq.PostID,
	})
	if createErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": createErr.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "新增成功",
	})

}

func (c *CmsApp) GetAllCommentByPostId(ctx *gin.Context) {
	var commentReq CommentReq
	if err := ctx.ShouldBindJSON(&commentReq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 记录日志
	logger.S().Debug("根据文章Id获取该文章的所有评论入参：", commentReq)

	// 根据id获取文章信息
	postDao := dao.NewPostDao(c.db)
	post, err := postDao.GetPostInfoById(commentReq.PostID)
	if err != nil {
		logger.S().Debug("GetAllCommentByPostId:根据文章Id获取文章信息异常", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if post == nil {
		logger.S().Debug("GetAllCommentByPostId:文章不存在，无法获取评论内容")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "文章不存在，无法获取评论内容"})
		return
	}

	commentDao := dao.NewCommentDao(c.db)
	comments, err := commentDao.GetAllCommentByPostId(post.ID)
	if err != nil {
		logger.S().Debug("GetAllCommentByPostId:根据文章Id获取所有评论异常", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": comments,
	})

}
