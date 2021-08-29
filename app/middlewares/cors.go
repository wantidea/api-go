package middlewares

import (
	"api-go/lib/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")

		var allowHeaders strings.Builder
		allowHeaders.WriteString("Content-Type,AccessToken,X-CSRF-Token,")
		allowHeaders.WriteString(jwt.UserTokenKey)
		allowHeaders.WriteString(",")
		allowHeaders.WriteString(jwt.AdminTokenKey)

		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", allowHeaders.String())
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT,PATCH")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
