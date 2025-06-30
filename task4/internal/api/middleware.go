package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"net/http"
	"task4/internal/common"
	util "task4/internal/utils"
)

// 中间件
// session鉴权

const SessionKey = "session_id"
const Authorization = "Authorization"

type SessionAuth struct {
}

type JWTTokenAuth struct {
}

func (s *SessionAuth) Auth(ctx *gin.Context) {
	sessionId := ctx.GetHeader(SessionKey)
	//  实现完整鉴权
	if sessionId == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": "Session为空或者失效"})
		return
	}

	authKey := util.GetAuthKey(sessionId)
	rdb := util.ConnRdb()
	loginTime, err := rdb.Get(ctx, authKey).Result()
	if err != nil && err != redis.Nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "Session鉴权失败1"})
		return
	}
	if loginTime == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": "Session鉴权失败2"})
		return
	}

	fmt.Println(sessionId)
	ctx.Next()
}

func (a JWTTokenAuth) AuthJWTToken(ctx *gin.Context) {
	//gin.HandlerFunc
	tokenString := ctx.GetHeader(Authorization)
	if tokenString == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未提供认证令牌"})
		return
	}

	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return []byte("task4_secret_key"), nil
		})

	if err != nil || !token.Valid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效令牌"})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	ctx.Set("id", claims["id"])
	ctx.Set("username", claims["username"])
	ctx.Set("isAuthentication", claims["isAuthentication"])
	ctx.Set("jwtClaims", claims)
	fmt.Println("JWTToken claims", claims)
	ctx.Next()
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 处理 panic 错误
		if err := recover(); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "服务器发生内部错误",
			})
			return
		}

		// 如果有错误，统一处理
		if len(c.Errors) > 0 {
			err := c.Errors[0].Err
			switch e := err.(type) {
			case *common.AppError:
				c.AbortWithStatusJSON(e.Code, gin.H{
					"code":    e.Code,
					"message": e.Message,
				})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":    http.StatusInternalServerError,
					"message": "服务器发生未知错误",
				})
			}
		}
	}
}
