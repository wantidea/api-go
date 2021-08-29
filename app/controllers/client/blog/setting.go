package blog

import (
	BlogModels "api-go/app/models/blog"
	"api-go/lib/file"
	"api-go/lib/orm"
	"api-go/lib/response"
	"github.com/gin-gonic/gin"
)

type Setting struct {
}

// List 获取所有设置
func (t *Setting) List(c *gin.Context) {
	appG := &response.Gin{C: c}

	var setting []map[string]interface{}
	orm.DB().Model(BlogModels.Setting{}).Find(&setting)

	result := map[string]interface{}{}
	for i := 0; i < len(setting); i++ {
		result[setting[i]["name"].(string)] = setting[i]["value"]

		if setting[i]["name"] == "logo" {
			result["logo_url"] = file.ImagePathToUrl(setting[i]["value"].(string))
		}
	}

	appG.SuccessResponse(response.CodeSuccess, result)
}
