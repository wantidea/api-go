package user

import (
	"api-go/lib/jwt"
	"api-go/lib/response"
	"github.com/gin-gonic/gin"
)

func InfoUser(c *gin.Context) {
	appG := &response.Gin{C: c}
	appG.SuccessResponse(response.CodeSuccess, jwt.UserInfo(c))
}
