package client

import (
	BlogControllers "api-go/app/controllers/client/blog"
	BlogArticleControllers "api-go/app/controllers/client/blog/article"
	"github.com/gin-gonic/gin"
)

// AddBlogRoute 添加 博客 路由
func AddBlogRoute(r *gin.Engine) {

	blogR := r.Group("/client/blog")

	// 控制器
	articleC := &BlogControllers.Article{}          // 文章
	categoryC := &BlogArticleControllers.Category{} // 文章分类
	settingC := &BlogControllers.Setting{}          // 设置

	// 文章
	articleR := blogR.Group("/article")
	{
		articleR.GET("/newList", articleC.NewList)         // 最新文章列表
		articleR.GET("/item", articleC.Item)               // 文章详情
		articleR.GET("/newQuitList", articleC.NewQuitList) // 侧边栏最新快捷导航
		articleR.GET("/hotQuitList", articleC.HotQuitList) // 热门文章

		// 文章分类
		categoryR := articleR.Group("/category")
		{
			categoryR.GET("/option", categoryC.Option)
			categoryR.GET("/articleCount", categoryC.ArticleCountList) // 分类文章数量
			categoryR.GET("/articleList", categoryC.ListArticle)       // 分类文章数量
			categoryR.GET("/item", categoryC.Item)
		}
	}

	// 设置
	settingR := blogR.Group("/setting")
	{
		settingR.GET("/list", settingC.List) // 获取所有设置
	}
}
