package service

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"net/http"
	"task4/internal/common/logger"
	"task4/internal/dao"
	"task4/internal/model"
	util "task4/internal/utils"
	"time"
)

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResp struct {
	SessionId string `json:"session_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Token     string `json:"token"`
}

func (c *CmsApp) Login(ctx *gin.Context) {
	var loginReq LoginReq
	loginInfo := &loginReq
	if err := ctx.ShouldBindJSON(loginInfo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var (
		userName = loginInfo.Username
		password = loginInfo.Password
	)
	logger.S().Info("正在进行登录的用户的userName为：%s", userName)
	userDao := dao.NewUserDao(c.db)
	userInfo, err := userDao.GetUserInfoByUserName(userName)
	if err != nil {
		logger.S().Debug("查询用户信息异常：%s", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := bcrypt.CompareHashAndPassword(
		[]byte(userInfo.Password),
		[]byte(password)); err != nil {
		logger.S().Debug("用户密码对比异常：%s", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "用户名或密码错误2"})
		return
	}

	sessionId, err := c.genterSessionId(context.Background(), userInfo.Username)
	if err != nil {
		logger.S().Debug("生成SessionId异常：%s", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.genterToken(userInfo)
	if err != nil {
		logger.S().Debug("生成JWToken异常：%s", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data": &LoginResp{
			SessionId: sessionId,
			Username:  userInfo.Username,
			Email:     userInfo.Email,
			Token:     token,
		},
	})
}

// 生成sessionId
func (c *CmsApp) genterSessionId(ctx context.Context, userName string) (string, error) {
	// 生成一个sessionId
	// key session_id:{userId} value: sessionId 过期时间：24h
	sessionId := uuid.New().String()
	sessionKey := util.GetSessionIdKey(userName)
	//fmt.Sprintf("session_id: %s", userName)
	err := c.rdb.Set(ctx, sessionKey, sessionId, time.Hour*24).Err()
	if err != nil {
		logger.S().Debug("genterSessionId rdb set", err)
		//panic(err)
		return sessionId, err
	}

	authKey := util.GetAuthKey(sessionId)
	// fmt.Sprintf("auth_key: %s", sessionId)
	err = c.rdb.Set(ctx, authKey, time.Now().Unix(), time.Hour*24).Err()
	if err != nil {
		logger.S().Debug("genterSessionId rdb set", err)
		//panic(err)
		return sessionId, err
	}

	logger.S().Debug("sessionKey:", sessionKey)
	return sessionId, nil
}

// 生成Token
func (c *CmsApp) genterToken(userInfo *model.User) (string, error) {
	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":               userInfo.ID,
		"username":         userInfo.Username,
		"isAuthentication": userInfo.IsAuthentication,
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	})

	// 必须调用SignedString方法获取最终的JWT字符串
	tokenString, err := token.SignedString([]byte("task4_secret_key"))
	if err != nil {
		logger.S().Debug("genterToken err", err)
		// 处理错误
		return "", err
	}

	logger.S().Debug("tokenString:", tokenString)
	return tokenString, nil
}
