package article

import (
	BlogModels "api-go/app/models/blog"
	BlogArticleModels "api-go/app/models/blog/article"
	BlogArticleRequests "api-go/app/requests/business/top/blog/article"
	UploadServices "api-go/app/services/upload"
	"api-go/lib/orm"
	"api-go/lib/page"
	"api-go/lib/request"
	"api-go/lib/response"
	"github.com/gin-gonic/gin"
	"time"
)

type Category struct {
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

// CountList 分类文章数量
func (t *Category) ArticleCountList(c *gin.Context) {
	appG := &response.Gin{C: c}

	article := &BlogModels.Article{}
	articleTable := article.TableName()
	category := &BlogArticleModels.Category{}
	categoryTable := category.TableName()

	var result []map[string]interface{}
	orm.DB().Raw("SELECT id, name,( SELECT count( id ) FROM " +
		articleTable +
		" AS t_b WHERE t_b.category_id = t_a.id AND t_b.deleted_at = 0) as total FROM " +
		categoryTable +
		" AS t_a WHERE t_a.deleted_at = 0").
		Find(&result)
	appG.SuccessResponse(response.CodeSuccess, result)
}

// ListArticle 分类文章列表
func (t *Category) ListArticle(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &BlogArticleRequests.ListArticleCategory{}

	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	var result []map[string]interface{}
	var count int64
	orm.DB().Model(&BlogModels.Article{}).
		Select(
			"id",
			"title",
			"editor_type",
			"banner_id",
			"created_at",
			"updated_at",
		).
		Where("category_id = ?", form.ID).
		Order("updated_at desc").
		Scopes(page.Paginate(c)).
		Find(&result)

	orm.DB().Model(&BlogModels.Article{}).
		Where("category_id = ?", form.ID).
		Count(&count)

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

		result[i]["banner_url"] = UploadServices.ImageUrl(result[i]["banner_id"].(int64))
	}

	data := map[string]interface{}{
		"data":  result,
		"count": count,
	}

	appG.SuccessResponse(response.CodeSuccess, data)
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
