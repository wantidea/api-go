package top

import (
	"github.com/gin-gonic/gin"
)

var firstPrefix = "/top" // 一级前缀

// AddTopRoute 添加 总站 路由
func AddTopRoute(r *gin.Engine) {
	AddBlogRoute(r)   // 博客
	AddUploadRoute(r) // 上传
	AddSystemRoute(r) // 系统
	AddUserRoute(r)   // 用户模块待重构
}
