package blog

import (
	BlogModels "api-go/app/models/blog"
	BlogArticleModels "api-go/app/models/blog/article"
	BlogRequests "api-go/app/requests/business/top/blog"
	CategoryServices "api-go/app/services/blog/article/category"
	UploadServices "api-go/app/services/upload"
	"api-go/lib/jwt"
	"api-go/lib/orm"
	"api-go/lib/page"
	"api-go/lib/request"
	"api-go/lib/response"
	"github.com/gin-gonic/gin"
	"time"
)

type Article struct {
}

// Add 添加文章
func (t *Article) Add(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &BlogRequests.AddArticle{}
	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	article := &BlogModels.Article{
		Title:         form.Title,
		EditorType:    form.EditorType,
		Content:       form.Content,
		ContentMd:     form.ContentMd,
		CategoryId:    form.CategoryId,
		BannerId:      form.BannerId,
		CreatedUserId: jwt.AdminId(c),
	}

	res := orm.DB().Create(article)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	} else {
		appG.SuccessResponse(response.CodeSuccess)
	}
}

// Del 删除文章
func (t *Article) Del(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &BlogRequests.DelArticle{}
	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	res := orm.DB().Delete(&BlogModels.Article{}, form.ID)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	} else {
		appG.SuccessResponse(response.CodeSuccess)
	}
}

// Upd 修改文章
func (t *Article) Upd(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &BlogRequests.UpdArticle{}
	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	article := &BlogModels.Article{}
	res := orm.DB().Where("id = ?", form.ID).First(article)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
		return
	}

	article.Title = form.Title
	article.Content = form.Content
	article.BannerId = form.BannerId
	article.CategoryId = form.CategoryId
	article.ContentMd = form.ContentMd

	res = orm.DB().Save(article)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	} else {
		appG.SuccessResponse(response.CodeSuccess)
	}
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

	appG.SuccessResponse(response.CodeSuccess, result)
}

// List 文章列表
func (t *Article) List(c *gin.Context) {
	type List struct {
		ID           int64  `json:"id"`
		Title        string `json:"title"`
		CategoryId   int64  `json:"category_id"`
		BannerId     int64  `json:"banner_id"`
		EditorType   int    `json:"editor_type"`
		LookTotal    int    `json:"look_total"`
		CreatedAt    int64  `json:"created_at"`
		UpdatedAt    int64  `json:"updated_at"`
		CategoryName string `json:"category_name" gorm:"column:category_name"`

		CreatedDate string `json:"created_date"`
		UpdatedDate string `json:"updated_date"`
		BannerUrl   string `json:"banner_url"`
	}

	appG := &response.Gin{C: c}
	var result []*List
	var count int64

	article := &BlogModels.Article{}
	category := &BlogArticleModels.Category{}

	articleTable := article.TableName()
	categoryTable := category.TableName()

	orm.DB().Table(articleTable).
		Select(
			articleTable+".id",
			articleTable+".title",
			articleTable+".category_id",
			articleTable+".banner_id",
			articleTable+".editor_type",
			articleTable+".look_total",
			articleTable+".created_at",
			articleTable+".updated_at",
			categoryTable+".name as category_name",
		).
		Joins("left join " + categoryTable + " on " + articleTable + ".category_id = " + categoryTable + ".id").
		Where(articleTable + ".deleted_at = 0 and " + categoryTable + ".deleted_at = 0").
		Order("updated_at desc").
		Scopes(page.Paginate(c)).
		Find(&result)

	orm.DB().Model(article).Count(&count)

	for i := 0; i < len(result); i++ {
		if result[i].CreatedAt != 0 {
			result[i].CreatedDate = time.Unix(result[i].CreatedAt, 0).Format("2006-01-02 15:04:05")
		} else {
			result[i].CreatedDate = "无"
		}

		if result[i].UpdatedAt != 0 {
			result[i].UpdatedDate = time.Unix(result[i].UpdatedAt, 0).Format("2006-01-02 15:04:05")
		} else {
			result[i].UpdatedDate = "无"
		}

		result[i].BannerUrl = UploadServices.ImageUrl(result[i].BannerId)
	}

	data := map[string]interface{}{
		"data":  result,
		"count": count,
	}

	appG.SuccessResponse(response.CodeSuccess, data)
}
