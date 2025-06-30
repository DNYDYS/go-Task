package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// GinLogger Gin 中间件
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		fields := []zap.Field{
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", time.Since(start)),
		}

		if len(c.Errors) > 0 {
			fields = append(fields, zap.String("errors", c.Errors.String()))
		}
		// 使用正确的访问方式
		loggedLogger := L().With(fields...)

		switch {
		case c.Writer.Status() >= 500:
			loggedLogger.Error("HTTP请求失败")
		case c.Writer.Status() >= 400:
			loggedLogger.Warn("HTTP请求错误")
		default:
			loggedLogger.Info("HTTP请求成功")
		}
	}
}

// GinRecovery 错误恢复中间件
func GinRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				S().With(
					zap.Any("error", err),
					zap.Stack("stack"),
				).Error("服务 panic 恢复")

				c.AbortWithStatusJSON(500, gin.H{
					"code":    500,
					"message": "服务器内部错误",
				})
			}
		}()
		c.Next()
	}
}
