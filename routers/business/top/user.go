package top

import (
	UserControllers "api-go/app/controllers/business/top/user"
	"api-go/app/middlewares"
	"github.com/gin-gonic/gin"
)

func AddUserRoute(r *gin.Engine) {
	userRouter := r.Group(firstPrefix + "/user")

	userRouter.Use(middlewares.AuthToken())
	{
		userRouter.GET("/info", UserControllers.InfoUser)
	}
}
