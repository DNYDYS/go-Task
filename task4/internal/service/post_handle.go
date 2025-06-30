package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	//"github.com/jinzhu/copier"
	"net/http"
	"task4/internal/dao"
	"task4/internal/model"
	util "task4/internal/utils"
)

type PostReq struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  uint   `json:"user_id"`
}

type PostResp struct {
	Message string `json:"message" binding:"required"`
}

func (c *CmsApp) CreatePost(ctx *gin.Context) {

	var newPostReq PostReq
	postReq := &newPostReq
	fmt.Println("新增文章接口 CreatePost入参:", *postReq)
	if err := ctx.ShouldBindJSON(postReq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 获取当前用户ID
	userId := util.GetUserIdFromCtx(ctx)

	if util.IsNilOrEmptyString(&postReq.Title) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "文章标题不能为空"})
		return
	}
	if util.IsNilOrEmptyString(&postReq.Content) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "文章内容不能为空"})
		return
	}

	userDao := dao.NewUserDao(c.db)
	userInfo, err := userDao.GetUserInfoByUserid(userId)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isAuthentication := userInfo.IsAuthentication

	if isAuthentication != "1" {
		// 0 未认证 1 已认证
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "用户未认证，请先认证"})
		return
	}

	postDao := dao.NewPostDao(c.db)
	if err := postDao.CreatPost(model.Post{
		Title:   postReq.Title,
		Content: postReq.Content,
		UserID:  userId,
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "文章新增成功",
		"data": &RegisterResp{
			Message: fmt.Sprintf("文章新增成功"),
		},
	})

}

func (c *CmsApp) GetAllPosts(ctx *gin.Context) {
	var newPostReq PostReq
	postReq := &newPostReq
	fmt.Println("<UNK> GetAllPosts<UNK>:", *postReq)
	if err := ctx.ShouldBindJSON(postReq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	postDao := dao.NewPostDao(c.db)
	var posts []model.Post
	if err := postDao.GetAllPosts(model.Post{
		Model:   gorm.Model{ID: postReq.Id}, // 手动设置 ID
		Title:   postReq.Title,
		Content: postReq.Content,
		UserID:  postReq.UserId,
	}).Find(&posts).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    posts,
	})
}

func (c *CmsApp) UpdatePostById(ctx *gin.Context) {
	var newPostReq PostReq
	postReq := &newPostReq
	fmt.Println("<UNK> GetPostById<UNK>:", *postReq)
	if err := ctx.ShouldBindJSON(postReq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	postDao := dao.NewPostDao(c.db)
	postInfo, err := postDao.GetPostInfoById(postReq.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 获取当前用户ID
	userId := util.GetUserIdFromCtx(ctx)
	if postInfo.UserID != userId {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "只能修改自己的文章"})
		return
	}

	postInfo.Title = postReq.Title
	postInfo.Content = postReq.Content
	updateerr := postDao.UpdatePostById(*postInfo)
	if updateerr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": updateerr.Error()})
		return
	}
	newPostInfo, err := postDao.GetPostInfoById(postReq.Id)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    newPostInfo,
	})

}

func (c *CmsApp) DeletePostById(ctx *gin.Context) {
	var newPostReq PostReq
	postReq := &newPostReq
	fmt.Println("<UNK> GetPostById<UNK>:", *postReq)
	if err := ctx.ShouldBindJSON(postReq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if postReq.Id == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "文章Id必传"})
		return
	}
	postDao := dao.NewPostDao(c.db)
	err := postDao.DeletePostById(postReq.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    "删除成功",
	})
}
