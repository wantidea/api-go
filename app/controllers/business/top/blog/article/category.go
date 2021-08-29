package article

import (
	BlogArticleModels "api-go/app/models/blog/article"
	BlogArticleRequests "api-go/app/requests/business/top/blog/article"
	"api-go/lib/jwt"
	"api-go/lib/orm"
	"api-go/lib/request"
	"api-go/lib/response"
	"github.com/gin-gonic/gin"
	"time"
)

type Category struct {
}

// Add 添加文章分类
func (t *Category) Add(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &BlogArticleRequests.AddCategory{}
	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	articleCategory := &BlogArticleModels.Category{
		Name:          form.Name,
		CreatedUserId: jwt.AdminId(c),
	}

	res := orm.DB().Create(articleCategory)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	} else {
		appG.SuccessResponse(response.CodeSuccess)
	}
}

// Del 删除文章分类
func (t *Category) Del(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &BlogArticleRequests.DelCategory{}
	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	res := orm.DB().Delete(&BlogArticleModels.Category{}, form.ID)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	} else {
		appG.SuccessResponse(response.CodeSuccess)
	}
}

// Upd 修改文章分类
func (t *Category) Upd(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &BlogArticleRequests.UpdCategory{}
	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	articleCategory := &BlogArticleModels.Category{}
	res := orm.DB().Where("id = ?", form.ID).First(articleCategory)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
		return
	}

	articleCategory.Name = form.Name

	res = orm.DB().Save(articleCategory)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	} else {
		appG.SuccessResponse(response.CodeSuccess)
	}
}

// Item 文章分类详情
func (t *Category) Item(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &BlogArticleRequests.ItemCategory{}

	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	result := map[string]interface{}{}
	res := orm.DB().Model(BlogArticleModels.Category{}).Where("id = ?", form.ID).First(&result)

	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
		return
	}

	if result["created_at"].(int) != 0 {
		result["created_date"] = time.Unix(int64(result["created_at"].(int)), 0).Format("2006-01-02")
	} else {
		result["created_date"] = "无"
	}

	appG.SuccessResponse(response.CodeSuccess, result)
}

// List 文章分类列表
func (t *Category) List(c *gin.Context) {
	appG := &response.Gin{C: c}
	var result []map[string]interface{}
	orm.DB().Model(&BlogArticleModels.Category{}).
		Select(
			"id",
			"name",
			"created_at",
			"updated_at",
		).
		Order("updated_at desc").
		Find(&result)

	for i := 0; i < len(result); i++ {
		if result[i]["created_at"].(int) != 0 {
			result[i]["created_date"] = time.Unix(int64(result[i]["created_at"].(int)), 0).Format("2006-01-02")
		} else {
			result[i]["created_date"] = "无"
		}

		if result[i]["updated_at"].(int) != 0 {
			result[i]["updated_date"] = time.Unix(int64(result[i]["updated_at"].(int)), 0).Format("2006-01-02")
		} else {
			result[i]["updated_date"] = "无"
		}
	}

	appG.SuccessResponse(response.CodeSuccess, &result)
}

// Option 文章分类选项
func (t *Category) Option(c *gin.Context) {
	appG := &response.Gin{C: c}
	var result []map[string]interface{}
	orm.DB().Model(&BlogArticleModels.Category{}).
		Select(
			"id",
			"name",
		).
		Find(&result)
	appG.SuccessResponse(response.CodeSuccess, &result)
}
