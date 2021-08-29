package system

import (
	SystemModels "api-go/app/models/system"
	SystemRequests "api-go/app/requests/business/top/system"
	"api-go/lib/jwt"
	"api-go/lib/orm"
	"api-go/lib/request"
	"api-go/lib/response"
	"github.com/gin-gonic/gin"
	"gorm.io/plugin/soft_delete"
	"time"
)

type Route struct {
}

func (t *Route) Add(c *gin.Context) {
	appG := &response.Gin{C: c}

	form := &SystemRequests.AddRoute{}
	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	route := &SystemModels.Route{
		Name:          form.Name,
		Uri:           form.Uri,
		Method:        form.Method,
		RoleList:      form.RoleList,
		IsAuth:        form.IsAuth,
		CreatedUserId: jwt.AdminId(c),
	}

	res := orm.DB().Create(route)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	} else {
		appG.SuccessResponse(response.CodeSuccess, route.ID)
	}
}

func (t *Route) Del(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &SystemRequests.DelRoute{}

	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	res := orm.DB().Delete(&SystemModels.Route{}, form.ID)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	} else {
		appG.SuccessResponse(response.CodeSuccess)
	}
}

func (t *Route) Upd(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &SystemRequests.UpdRoute{}

	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	route := &SystemModels.Route{}
	res := orm.DB().Where("id = ?", form.ID).First(route)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
		return
	}

	route.Name = form.Name
	route.Uri = form.Uri
	route.Method = form.Method
	route.RoleList = form.RoleList
	route.IsAuth = form.IsAuth

	res = orm.DB().Save(route)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	} else {
		appG.SuccessResponse(response.CodeSuccess)
	}
}

func (t *Route) Item(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &SystemRequests.ItemRoute{}

	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	result := map[string]interface{}{}
	res := orm.DB().Model(&SystemModels.Route{}).Where("id = ?", form.ID).First(&result)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError, &result)
		return
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

	if result["deleted_at"].(soft_delete.DeletedAt) != 0 {
		result["deleted_date"] = time.Unix(int64(result["deleted_at"].(soft_delete.DeletedAt)), 0).Format("2006-01-02")
	} else {
		result["deleted_date"] = "无"
	}

	appG.SuccessResponse(response.CodeSuccess, &result)
}

// List 查询角色列表
func (t *Route) List(c *gin.Context) {
	appG := &response.Gin{C: c}
	var result []map[string]interface{}
	orm.DB().Model(&SystemModels.Route{}).Find(&result)

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

		if result[i]["deleted_at"].(soft_delete.DeletedAt) != 0 {
			result[i]["deleted_date"] = time.Unix(int64(result[i]["deleted_at"].(soft_delete.DeletedAt)), 0).Format("2006-01-02")
		} else {
			result[i]["deleted_date"] = "无"
		}
	}

	appG.SuccessResponse(response.CodeSuccess, &result)
}
