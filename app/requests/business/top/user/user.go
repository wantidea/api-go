package user

// AddUser 添加用户表单
type AddUser struct {
	Name     string `form:"name" json:"name" fieldName:"用户名" binding:"required,alphanum,min=3,max=30"`
	Password string `form:"password" json:"password" fieldName:"密码" binding:"required,min=8,max=100"`
	RoleId   int64  `form:"role_id" json:"role_id" fieldName:"角色" binding:"required"`
}

// DelUser 删除用户表单
type DelUser struct {
	ID string `form:"id" json:"id" fieldName:"ID" binding:"required"`
}

// UpdUser 修改用户表单
type UpdUser struct {
	ID       string `form:"id" json:"id" fieldName:"ID" binding:"required"`
	Name     string `form:"name" json:"name" fieldName:"用户名" binding:"required,alphanum,min=3,max=30"`
	Password string `form:"password" json:"password" fieldName:"密码" binding:"required,min=8,max=100"`
	RoleId   string `form:"role_id" json:"role_id" fieldName:"角色" binding:"required"`
}

// ItemUser 查询用户表单
type ItemUser struct {
	ID string `form:"id" json:"id" fieldName:"ID" binding:"required"`
}

// ListUser 查询用户表单
type ListUser struct {
	Name string `form:"name" json:"name" fieldName:"用户名"`
}

// NameUser 用户名查询用户表单
type NameUser struct {
	Name string `form:"name" json:"name" fieldName:"用户名" binding:"required,alphanum,min=3,max=30"`
}
