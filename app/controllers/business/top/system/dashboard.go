package system

import (
	ArticleServices "api-go/app/services/blog/article"
	"api-go/app/services/system"
	"api-go/lib/response"
	"github.com/gin-gonic/gin"
)

type Dashboard struct {
}

// List 仪表台信息列表
func (t *Dashboard) List(c *gin.Context) {
	appG := &response.Gin{C: c}
	article := &ArticleServices.Article{}

	type List struct {
		AdminTotal   int64 `json:"admin_total"`
		RoleTotal    int64 `json:"role_total"`
		ArticleTotal int64 `json:"article_total"`
	}

	list := &List{}
	list.AdminTotal = system.AdminTotalOnRedis()
	list.RoleTotal = system.RoleTotalOnRedis()
	list.ArticleTotal = article.TotalOnRedis()

	appG.SuccessResponse(response.CodeSuccess, list)
}
