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

type Role struct {
}

// Add 添加角色
func (t *Role) Add(c *gin.Context) {
	appG := &response.Gin{C: c}

	form := &SystemRequests.AddRole{}
	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	role := &SystemModels.Role{
		Name:          form.Name,
		Description:   form.Description,
		CreatedUserId: jwt.AdminId(c),
	}

	err := role.CreateCheckName()
	if err != nil {
		appG.ErrorMsgResponse(err.Error())
		return
	}

	res := orm.DB().Create(role)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	} else {
		appG.SuccessResponse(response.CodeSuccess, role.ID)
	}
}

// Del 删除角色
func (t *Role) Del(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &SystemRequests.DelRole{}

	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	res := orm.DB().Delete(&SystemModels.Role{}, form.ID)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	} else {
		appG.SuccessResponse(response.CodeSuccess)
	}
}

// Upd 修改角色
func (t *Role) Upd(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &SystemRequests.UpdRole{}

	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	role := &SystemModels.Role{}
	res := orm.DB().Where("id = ?", form.ID).First(role)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
		return
	}

	role.Name = form.Name
	err := role.UpdateCheckName()
	if err != nil {
		appG.ErrorMsgResponse(err.Error())
		return
	}
	role.Description = form.Description

	res = orm.DB().Save(role)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	} else {
		appG.SuccessResponse(response.CodeSuccess)
	}
}

// Item 查询角色详情
func (t *Role) Item(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &SystemRequests.ItemRole{}

	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	result := map[string]interface{}{}
	res := orm.DB().Model(&SystemModels.Role{}).Where("id = ?", form.ID).First(&result)
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
func (t *Role) List(c *gin.Context) {
	appG := &response.Gin{C: c}
	var result []map[string]interface{}
	orm.DB().Model(&SystemModels.Role{}).
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

// Option
func (t *Role) Option(c *gin.Context) {
	appG := &response.Gin{C: c}
	var result []map[string]interface{}
	orm.DB().Model(&SystemModels.Role{}).Select("id,name").Find(&result)
	appG.SuccessResponse(response.CodeSuccess, &result)
}
