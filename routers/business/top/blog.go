package top

import (
	BlogControllers "api-go/app/controllers/business/top/blog"
	BlogSettingControllers "api-go/app/controllers/business/top/blog"
	BlogArticleControllers "api-go/app/controllers/business/top/blog/article"
	"api-go/app/middlewares"
	"github.com/gin-gonic/gin"
)

func AddBlogRoute(r *gin.Engine) {
	blogR := r.Group(firstPrefix + "/blog")

	// 控制器
	articleC := &BlogControllers.Article{}
	categoryC := &BlogArticleControllers.Category{}
	settingC := &BlogSettingControllers.Setting{}

	// 文章
	articleR := blogR.Group("/article")
	// 需认证
	articleR.Use(middlewares.AdminAuthToken())
	{
		articleR.POST("/add", articleC.Add)
		articleR.DELETE("/del", articleC.Del)
		articleR.PUT("/upd", articleC.Upd)
		articleR.GET("/item", articleC.Item)
		articleR.GET("/list", articleC.List)

		// 文章分类
		categoryR := articleR.Group("/category")
		{
			categoryR.POST("/add", categoryC.Add)
			categoryR.DELETE("/del", categoryC.Del)
			categoryR.PUT("/upd", categoryC.Upd)
			categoryR.GET("/item", categoryC.Item)
			categoryR.GET("/list", categoryC.List)
			categoryR.GET("/option", categoryC.Option)
		}
	}

	// 设置
	setting := blogR.Group("/setting")
	// 需认证
	setting.Use(middlewares.AdminAuthToken())
	{
		setting.PATCH("/", settingC.Patch)  // 通过 name 修改相应设置
		setting.GET("/name", settingC.Name) // 通过 name 读取相应设置
	}

}
