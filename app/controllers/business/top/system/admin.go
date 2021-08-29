package system

import (
	SystemModels "api-go/app/models/system"
	SystemRequests "api-go/app/requests/business/top/system"
	UploadServices "api-go/app/services/upload"
	"api-go/lib/jwt"
	"api-go/lib/md5"
	"api-go/lib/orm"
	"api-go/lib/request"
	"api-go/lib/response"
	"github.com/gin-gonic/gin"
	"gorm.io/plugin/soft_delete"
	"time"
)

type Admin struct {
}

// Add添加管理员
func (t *Admin) Add(c *gin.Context) {
	appG := &response.Gin{C: c}

	form := &SystemRequests.AddAdmin{}
	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	admin := &SystemModels.Admin{
		Name:          form.Name,
		Password:      md5.ToMD5Salt(form.Password),
		RoleId:        form.RoleId,
		CreatedUserId: jwt.AdminId(c),
	}

	err := admin.CreateCheckName()
	if err != nil {
		appG.ErrorMsgResponse(err.Error())
		return
	}

	res := orm.DB().Create(admin)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeErrorUserAdd)
	} else {
		appG.SuccessResponse(response.CodeSuccessUserAdd, admin.ID)
	}
}

// Del删除管理员
func (t *Admin) Del(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &SystemRequests.DelAdmin{}

	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	res := orm.DB().Delete(&SystemModels.Admin{}, form.ID)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	} else {
		appG.SuccessResponse(response.CodeSuccess)
	}
}

// Upd修改管理员
func (t *Admin) Upd(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &SystemRequests.UpdAdmin{}

	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	admin := &SystemModels.Admin{}
	res := orm.DB().Where("id = ?", form.ID).First(admin)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
		return
	}

	admin.Name = form.Name
	admin.RoleId = form.RoleId
	admin.AvatarId = form.AvatarId
	err := admin.UpdateCheckName()
	if err != nil {
		appG.ErrorMsgResponse(err.Error())
		return
	}

	res = orm.DB().Save(admin)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	} else {
		appG.SuccessResponse(response.CodeSuccess)
	}

}

// Item管理员列表
func (t *Admin) Item(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &SystemRequests.ItemAdmin{}

	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	result := map[string]interface{}{}
	res := orm.DB().Model(&SystemModels.Admin{}).Where("id = ?", form.ID).First(&result)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError, &result)
		return
	}

	result["avatar_url"] = UploadServices.ImageUrl(result["avatar_id"].(int64))

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

// List管理员列表
func (t *Admin) List(c *gin.Context) {
	appG := &response.Gin{C: c}
	var result []map[string]interface{}

	admin := &SystemModels.Admin{}
	role := &SystemModels.Role{}

	adminTable := admin.TableName()
	roleTable := role.TableName()

	orm.DB().Table(adminTable).
		Select(
			adminTable+".id",
			adminTable+".name",
			adminTable+".role_id",
			adminTable+".created_at",
			adminTable+".updated_at",
			roleTable+".name as role_name",
		).
		Joins("left join " + roleTable + " on " + adminTable + ".role_id = " + roleTable + ".id and " + roleTable + ".deleted_at = 0").
		Where(adminTable + ".deleted_at = 0").
		Find(&result)

	for i := 0; i < len(result); i++ {
		if result[i]["created_at"].(int32) != 0 {
			result[i]["created_date"] = time.Unix(int64(result[i]["created_at"].(int32)), 0).Format("2006-01-02")
		} else {
			result[i]["created_date"] = "无"
		}

		if result[i]["updated_at"].(int32) != 0 {
			result[i]["updated_date"] = time.Unix(int64(result[i]["updated_at"].(int32)), 0).Format("2006-01-02")
		} else {
			result[i]["updated_date"] = "无"
		}
	}

	appG.SuccessResponse(response.CodeSuccess, &result)
}

// Login管理员登录
func (t *Admin) Login(c *gin.Context) {
	var err error
	appG := &response.Gin{C: c}
	form := &SystemRequests.AuthAdmin{}

	err = request.CheckForm(c, appG, &form)
	if err != nil {
		return
	}

	token := ""
	admin := &SystemModels.Admin{
		Name:     form.Name,
		Password: form.Password,
	}
	if admin.Auth() {
		token, err = jwt.GenerateAdminToken(admin)
		if err != nil {
			appG.ErrorMsgResponse("生成令牌失败")
			return
		}
	} else {
		appG.ErrorMsgResponse("登录失败，请检查用户名密码")
		return
	}

	result := map[string]interface{}{
		"token": token,
	}
	appG.SuccessResponse(response.CodeSuccessAuth, result)
}

// Info管理员信息
func (t *Admin) Info(c *gin.Context) {
	appG := &response.Gin{C: c}
	appG.SuccessResponse(response.CodeSuccess, jwt.AdminInfo(c))
}

// RePwd重置管理员密码
func (t *Admin) RePwd(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &SystemRequests.RePwdAdmin{}

	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	admin := &SystemModels.Admin{}
	res := orm.DB().Where("id = ?", form.ID).First(admin)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
		return
	}

	admin.Password = md5.ToMD5Salt(form.Password)
	res = orm.DB().Save(admin)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	} else {
		appG.SuccessResponse(response.CodeSuccess)
	}
}
