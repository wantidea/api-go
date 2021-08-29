package blog

import (
	BlogModels "api-go/app/models/blog"
	BlogRequests "api-go/app/requests/business/top/blog"
	"api-go/lib/file"
	"api-go/lib/orm"
	"api-go/lib/request"
	"api-go/lib/response"
	"github.com/gin-gonic/gin"
)

type Setting struct {
}

// Patch 修改设置
func (t *Setting) Patch(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &BlogRequests.UpdSetting{}
	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	setting := &BlogModels.Setting{}
	res := orm.DB().Where("name = ?", form.Name).First(setting)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
		return
	}

	setting.Name = form.Name
	setting.Value = form.Value

	res = orm.DB().Save(setting)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	} else {
		appG.SuccessResponse(response.CodeSuccess)
	}
}

// Name 查询设置
func (t *Setting) Name(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &BlogRequests.NameSetting{}
	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	result := map[string]interface{}{}
	res := orm.DB().Model(BlogModels.Setting{}).Where("name = ?", form.Name).First(&result)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
		return
	}

	if form.Name == "logo" {
		result["logo_url"] = file.ImagePathToUrl(result["value"].(string))
	}

	appG.SuccessResponse(response.CodeSuccess, result)
}
