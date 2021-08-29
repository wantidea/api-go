package system

import (
	SystemModels "api-go/app/models/system"
	SystemRequests "api-go/app/requests/business/top/system"
	SystemServices "api-go/app/services/system"
	"api-go/lib/jwt"
	"api-go/lib/orm"
	"api-go/lib/request"
	"api-go/lib/response"
	"github.com/gin-gonic/gin"
)

type Menu struct {
}

// Add
func (t *Menu) Add(c *gin.Context) {
	appG := &response.Gin{C: c}

	form := &SystemRequests.AddMenu{}
	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	level := 0
	if form.ParentId != 0 {
		parentMenu := &SystemModels.Menu{}
		res := orm.DB().Model(parentMenu).
			Select("id", "level").
			Where("id = ?", form.ParentId).
			First(&parentMenu)
		if res.Error != nil {
			appG.ErrorResponse(response.CodeError)
			return
		}
		level = parentMenu.Level + 1
	}

	menu := &SystemModels.Menu{
		ParentId:      form.ParentId,
		Level:         level,
		Name:          form.Name,
		Title:         form.Title,
		Path:          form.Path,
		Redirect:      form.Redirect,
		Component:     form.Component,
		Icon:          form.Icon,
		Sort:          form.Sort,
		RoleList:      form.RoleList,
		IsHidden:      form.IsHidden,
		IsAuth:        form.IsAuth,
		CreatedUserId: jwt.AdminId(c),
	}

	res := orm.DB().Create(menu)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	} else {
		_ = SystemServices.AddRefreshTreeMenuProducer()
		appG.SuccessResponse(response.CodeSuccess, menu.ID)
	}
}

// Del
func (t *Menu) Del(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &SystemRequests.DelMenu{}

	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	res := orm.DB().Delete(&SystemModels.Menu{}, form.ID)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	} else {
		_ = SystemServices.AddRefreshTreeMenuProducer()
		appG.SuccessResponse(response.CodeSuccess)
	}
}

// Upd
func (t *Menu) Upd(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &SystemRequests.UpdMenu{}

	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	menu := &SystemModels.Menu{}
	res := orm.DB().Where("id = ?", form.ID).First(menu)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
		return
	}

	level := 0
	if form.ParentId != 0 {
		parentMenu := &SystemModels.Menu{}
		res := orm.DB().Model(parentMenu).
			Where("id = ?", form.ParentId).
			Select("id", "level").
			First(&parentMenu)
		if res.Error != nil {
			appG.ErrorResponse(response.CodeError)
			return
		}
		level = parentMenu.Level + 1
	}

	menu.ParentId = form.ParentId
	menu.Level = level
	menu.Name = form.Name
	menu.Title = form.Title
	menu.Path = form.Path
	menu.Redirect = form.Redirect
	menu.Component = form.Component
	menu.Icon = form.Icon
	menu.Sort = form.Sort
	menu.RoleList = form.RoleList
	menu.IsHidden = form.IsHidden
	menu.IsAuth = form.IsAuth

	res = orm.DB().Save(menu)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	} else {
		_ = SystemServices.AddRefreshTreeMenuProducer()
		appG.SuccessResponse(response.CodeSuccess)
	}
}

// Item 管理员列表
func (t *Menu) Item(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &SystemRequests.ItemMenu{}

	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	result := map[string]interface{}{}
	res := orm.DB().Model(&SystemModels.Menu{}).Where("id = ?", form.ID).Select(
		"parent_id",
		"level",
		"name",
		"title",
		"path",
		"redirect",
		"component",
		"icon",
		"sort",
		"role_list",
		"is_hidden",
		"is_auth",
	).First(&result)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError, &result)
		return
	}

	appG.SuccessResponse(response.CodeSuccess, &result)
}

// List 管理员列表
func (t *Menu) List(c *gin.Context) {
	appG := &response.Gin{C: c}
	var result []map[string]interface{}
	orm.DB().Table("system_menu as sm_a").
		Select(
			"sm_a.id",
			"sm_a.title",
			"sm_a.path",
			"sm_a.redirect",
			"sm_a.component",
			"sm_a.icon",
			"sm_a.sort",
			"sm_a.is_hidden",
			"sm_a.is_auth",
			"sm_b.title as parent_title",
		).
		Joins("left join system_menu as sm_b ON sm_a.parent_id = sm_b.id").
		Order("sm_a.sort asc").
		Find(&result)

	for i := 0; i < len(result); i++ {
		if result[i]["parent_title"] == nil {
			result[i]["parent_title"] = "无"
		}

		if result[i]["is_hidden"].(int8) == 0 {
			result[i]["hidden"] = "显示"
		} else {
			result[i]["hidden"] = "隐藏"
		}

		if result[i]["is_auth"].(int8) == 0 {
			result[i]["auth"] = "不启用"
		} else {
			result[i]["auth"] = "启用"
		}
	}

	appG.SuccessResponse(response.CodeSuccess, result)
}

// Async 动态菜单
func (t *Menu) Async(c *gin.Context) {
	appG := &response.Gin{C: c}
	admin := jwt.AdminInfo(c)

	result, _ := SystemServices.TreeMenuOnAuthRedisByRoleId(admin.RoleId)
	appG.SuccessResponse(response.CodeSuccess, result)
}

// Option 菜单选项
func (t *Menu) Option(c *gin.Context) {
	appG := &response.Gin{C: c}
	var result []map[string]interface{}
	orm.DB().Model(&SystemModels.Menu{}).Select(
		"id",
		"title",
	).Where("is_hidden = 0").Find(&result)

	appG.SuccessResponse(response.CodeSuccess, result)
}

// TreeList 树形列表菜单
func (t *Menu) TreeList(c *gin.Context) {
	appG := &response.Gin{C: c}
	var result []*SystemServices.MenuTree
	result = SystemServices.TreeMenu()
	appG.SuccessResponse(response.CodeSuccess, result)
}

// PatchIsAuth 启用或取消认证
func (t *Menu) PatchIsAuth(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &SystemRequests.PatchIsAuthMenu{}
	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	menu := &SystemModels.Menu{}
	res := orm.DB().Where("id = ?", form.ID).First(menu)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	}

	isAuth := 0
	if form.IsAuth == 1 {
		isAuth = 1
	}
	menu.IsAuth = isAuth
	res = orm.DB().Save(menu)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
		return
	}

	_ = SystemServices.AddRefreshTreeMenuProducer()
	appG.SuccessResponse(response.CodeSuccess)
}

// PatchPath 修改菜单路径
func (t *Menu) PatchPath(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &SystemRequests.PatchPathMenu{}
	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	menu := &SystemModels.Menu{}
	res := orm.DB().Where("id = ?", form.ID).First(menu)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	}

	menu.Path = form.Path
	res = orm.DB().Save(menu)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
		return
	}

	_ = SystemServices.AddRefreshTreeMenuProducer()
	appG.SuccessResponse(response.CodeSuccess)
}

// PatchRedirect 修改菜单重定向路径
func (t *Menu) PatchRedirect(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &SystemRequests.PatchRedirectMenu{}
	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	menu := &SystemModels.Menu{}
	res := orm.DB().Where("id = ?", form.ID).First(menu)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	}

	menu.Redirect = form.Redirect
	res = orm.DB().Save(menu)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
		return
	}

	_ = SystemServices.AddRefreshTreeMenuProducer()
	appG.SuccessResponse(response.CodeSuccess)
}

// PatchIcon 修改菜单图标
func (t *Menu) PatchIcon(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &SystemRequests.PatchIconMenu{}
	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	menu := &SystemModels.Menu{}
	res := orm.DB().Where("id = ?", form.ID).First(menu)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	}

	menu.Icon = form.Icon
	res = orm.DB().Save(menu)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
		return
	}

	_ = SystemServices.AddRefreshTreeMenuProducer()
	appG.SuccessResponse(response.CodeSuccess)
}

// Move 移动菜单
func (t *Menu) Move(c *gin.Context) {
	appG := &response.Gin{C: c}
	form := &SystemRequests.MoveAuthMenu{}
	checkError := request.CheckForm(c, appG, form)
	if checkError != nil {
		return
	}

	menu := &SystemModels.Menu{}
	endMenu := &SystemModels.Menu{}
	menuRes := orm.DB().Where("id = ?", form.ID).First(menu)
	endMenuRes := orm.DB().Where("id = ?", form.EndId).First(endMenu)
	if menuRes.Error != nil || endMenuRes.Error != nil {
		appG.ErrorResponse(response.CodeError)
	}

	switch form.Position {
	case "before":
		menu.ParentId = endMenu.ParentId
		menu.Level = endMenu.Level
		menu.Sort = 0
		if endMenu.Sort > 0 {
			menu.Sort = endMenu.Sort - 1
		}
	case "after":
		menu.ParentId = endMenu.ParentId
		menu.Level = endMenu.Level
		menu.Sort = endMenu.Sort + 1
	case "inner":
		menu.ParentId = endMenu.ID
		menu.Level = endMenu.Level + 1
		menu.Sort = endMenu.Sort + 1
	default:
		appG.ErrorResponse(response.CodeError)
		return
	}

	res := orm.DB().Save(menu)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeError)
	}

	_ = SystemServices.AddRefreshTreeMenuProducer()
	appG.SuccessResponse(response.CodeSuccess)
}
