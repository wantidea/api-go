package top

import (
	UploadControllers "api-go/app/controllers/business/top/upload"
	"api-go/app/middlewares"
	"github.com/gin-gonic/gin"
)

// AddUploadRoute 添加 上传 路由
func AddUploadRoute(r *gin.Engine) {

	uploadR := r.Group(firstPrefix + "/upload")

	// 控制器
	imageC := &UploadControllers.Image{}

	// 图片
	imageR := uploadR.Group("/image")
	{
		imageR.POST("/secretUpload", imageC.SecretUpload)
	}
	// 需认证
	imageR.Use(middlewares.AdminAuthToken())
	{
		imageR.POST("/upload", imageC.Upload)
		imageR.GET("/list", imageC.List)
		imageR.PUT("/allUrl", imageC.UpdAllUrl)
	}
}
