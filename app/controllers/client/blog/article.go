package blog

import (
	BlogModels "api-go/app/models/blog"
	BlogRequests "api-go/app/requests/client/blog"
	CategoryServices "api-go/app/services/blog/article/category"
	UploadServices "api-go/app/services/upload"
	"api-go/lib/orm"
	"api-go/lib/page"
	"api-go/lib/request"
	"api-go/lib/response"
	"github.com/gin-gonic/gin"
	"time"
)

type Article struct {
}

// NewList 最新文章列表
func (t *Article) NewList(c *gin.Context) {
	appG := &response.Gin{C: c}
	var result []map[string]interface{}
	var count int64
	orm.DB().Model(&BlogModels.Article{}).
		Select(
			"id",
			"title",
			"banner_id",
			"created_at",
			"updated_at",
		).
		Order("created_at desc").
		Scopes(page.Paginate(c)).
		Find(&result)

	orm.DB().Model(&BlogModels.Article{}).Count(&count)

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

// Item 文章详情
func (t *Article) Item(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &BlogRequests.ItemArticle{}

	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	article := &BlogModels.Article{}
	res := orm.DB().Model(BlogModels.Article{}).Where("id = ?", form.ID).First(&article)

	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
		return
	}

	result := map[string]interface{}{
		"id":              article.ID,
		"title":           article.Title,
		"category_id":     article.CategoryId,
		"category_name":   CategoryServices.NameById(article.CategoryId),
		"banner_id":       article.BannerId,
		"editor_type":     article.BannerId,
		"content":         article.Content,
		"content_md":      article.ContentMd,
		"look_total":      article.LookTotal,
		"created_user_id": article.CreatedUserId,
		"created_at":      article.CreatedAt,
		"updated_at":      article.UpdatedAt,
		"banner_url":      UploadServices.ImageUrl(article.BannerId),
		"created_date":    "",
		"updated_date":    "",
	}

	if result["created_at"].(int) != 0 {
		result["created_date"] = time.Unix(int64(result["created_at"].(int)), 0).Format("2006-01-02")
	} else {
		result["created_date"] = "无"
	}
	if result["updated_at"].(int) != 0 {
		result["updated_date"] = time.Unix(int64(result["updated_at"].(int)), 0).Format("2006-01-02")
	} else {
		result["updated_date"] = "无"
	}

	orm.DB().Model(&article).Update("look_total", article.LookTotal+1)

	appG.SuccessResponse(response.CodeSuccess, result)
}

// NewQuitList 侧边栏最新快捷导航
func (t *Article) NewQuitList(c *gin.Context) {
	// 快捷导航数量
	limit := 5
	appG := &response.Gin{C: c}
	var result []map[string]interface{}
	orm.DB().Model(&BlogModels.Article{}).
		Select(
			"id",
			"title",
			"banner_id",
			"created_at",
		).
		Order("updated_at desc").
		Limit(limit).
		Find(&result)

	for i := 0; i < len(result); i++ {
		if result[i]["created_at"].(int) != 0 {
			result[i]["created_date"] = time.Unix(int64(result[i]["created_at"].(int)), 0).Format("2006-01-02")
		} else {
			result[i]["created_date"] = "无"
		}

		result[i]["banner_url"] = UploadServices.ImageUrl(result[i]["banner_id"].(int64))
	}

	appG.SuccessResponse(response.CodeSuccess, &result)
}

// HotQuitList 热门文章
func (t *Article) HotQuitList(c *gin.Context) {
	// 快捷导航数量
	limit := 5
	appG := &response.Gin{C: c}
	var result []map[string]interface{}
	orm.DB().Model(&BlogModels.Article{}).
		Select(
			"id",
			"title",
			"look_total",
		).
		Order("look_total desc").
		Limit(limit).
		Find(&result)
	appG.SuccessResponse(response.CodeSuccess, &result)
}
