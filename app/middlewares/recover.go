package middlewares

import (
	"api-go/lib/logger"
	"api-go/lib/response"
	"github.com/gin-gonic/gin"
)

// 捕获全局错误
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				appG := response.Gin{C: c}
				logger.Error(err)
				appG.ErrorResponse(response.CodeErrorSystem)
				c.Abort()
			}
		}()
		c.Next()
	}
}
