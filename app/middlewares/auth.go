package middlewares

import (
	"api-go/lib/jwt"
	"api-go/lib/response"
	"github.com/gin-gonic/gin"
	"time"
)

// AuthToken token 认证
func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := &response.Gin{C: c}
		var code int
		var data interface{}

		code = response.CodeSuccess
		token := c.Request.Header.Get("token")
		if token == "" {
			code = response.CodeErrorAuthCheckTokenNull
		} else {
			claims, err := jwt.ParseToken(token)
			if err != nil {
				code = response.CodeErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = response.CodeErrorAuthCheckTokenTimeout
			}
		}

		if code != response.CodeSuccess {
			appG.ErrorResponse(code, data)
			c.Abort()
			return
		}

		c.Next()
	}
}

// AdminAuthToken token 认证
func AdminAuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := &response.Gin{C: c}
		var code int
		var data interface{}

		code = response.CodeSuccess
		token := c.Request.Header.Get(jwt.AdminTokenKey)
		if token == "" {
			code = response.CodeErrorAuthCheckTokenNull
		} else {
			claims, err := jwt.ParseAdminToken(token)
			if err != nil {
				code = response.CodeErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = response.CodeErrorAuthCheckTokenTimeout
			}
		}

		if code != response.CodeSuccess {
			appG.ErrorResponse(code, data)
			c.Abort()
			return
		}

		c.Next()
	}
}
