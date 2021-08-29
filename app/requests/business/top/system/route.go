package system

// AddRoute
type AddRoute struct {
	Name     string `form:"name" json:"name" fieldName:"路由名称" binding:"required"`
	Uri      string `form:"uri" json:"uri" fieldName:"路由接口" binding:"required"`
	Method   string `form:"method" json:"method" fieldName:"请求方法" binding:"required"`
	RoleList string `form:"role_list" json:"role_list" fieldName:"角色组" binding:"required"`
	IsAuth   int8   `form:"is_auth" json:"is_auth" fieldName:"启用认证" binding:"required"`
}

// DelRoute
type DelRoute struct {
	ID int64 `form:"id" json:"id" fieldName:"ID" binding:"required"`
}

// UpdRoute
type UpdRoute struct {
	ID       int64  `form:"id" json:"id" fieldName:"ID" binding:"required"`
	Name     string `form:"name" json:"name" fieldName:"路由名称" binding:"required"`
	Uri      string `form:"uri" json:"uri" fieldName:"路由接口" binding:"required"`
	Method   string `form:"method" json:"method" fieldName:"请求方法" binding:"required"`
	RoleList string `form:"role_list" json:"role_list" fieldName:"角色组" binding:"required"`
	IsAuth   int8   `form:"is_auth" json:"is_auth" fieldName:"启用认证" binding:"required"`
}

// ItemRoute
type ItemRoute struct {
	ID int64 `form:"id" json:"id" fieldName:"ID" binding:"required"`
}

// ListRoute
type ListRoute struct{}
