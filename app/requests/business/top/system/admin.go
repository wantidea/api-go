package system

// AuthAdmin 添加管理员
type AuthAdmin struct {
	Name     string `form:"name" json:"name" fieldName:"用户名" binding:"required,alphanum,min=3,max=30"`
	Password string `form:"password" fieldName:"密码" json:"password" binding:"required,min=8,max=100"`
}

// AddAdmin 添加管理员
type AddAdmin struct {
	Name     string `form:"name" json:"name" fieldName:"管理员名称" binding:"required,alphanum,min=3,max=30"`
	Password string `form:"password" json:"password" fieldName:"密码" binding:"required,min=8,max=100"`
	RoleId   int64  `form:"role_id" json:"role_id" fieldName:"角色" binding:"required"`
}

// DelAdmin 删除管理员
type DelAdmin struct {
	ID int64 `form:"id" json:"id" fieldName:"ID" binding:"required"`
}

// UpdAdmin 修改管理员
type UpdAdmin struct {
	ID       int64  `form:"id" json:"id" fieldName:"ID" binding:"required"`
	Name     string `form:"name" json:"name" fieldName:"管理员名称" binding:"required,alphanum,min=3,max=30"`
	RoleId   int64  `form:"role_id" json:"role_id" fieldName:"角色" binding:"required"`
	AvatarId int64  `form:"avatar_id" json:"avatar_id" fieldName:"头像" binding:""`
}

// ItemAdmin 查询管理员
type ItemAdmin struct {
	ID int64 `form:"id" json:"id" fieldName:"ID" binding:"required"`
}

// ListAdmin 查询管理员
type ListAdmin struct {
	Name string `form:"name" json:"name" fieldName:"管理员名称"`
}

// NameAdmin 管理员名称查询管理员
type NameAdmin struct {
	Name string `form:"name" json:"name" fieldName:"管理员名称" binding:"required,alphanum,min=3,max=30"`
}

// RePwdAdmin 重置管理员密码
type RePwdAdmin struct {
	ID       int64  `form:"id" json:"id" fieldName:"ID" binding:"required"`
	Password string `form:"password" json:"password" fieldName:"密码" binding:"required,min=8,max=100"`
}
