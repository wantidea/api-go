package system

// AddRole 添加角色表单
type AddRole struct {
	Name        string `form:"name" json:"name" fieldName:"角色名" binding:"required,min=3,max=30"`
	Description string `form:"description" json:"description" fieldName:"描述" binding:"required,min=2,max=100"`
}

// DelRole 删除角色表单
type DelRole struct {
	ID int64 `form:"id" json:"id" fieldName:"ID" binding:"required"`
}

// UpdRole 修改角色表单
type UpdRole struct {
	ID          int64  `form:"id" json:"id" fieldName:"ID" binding:"required"`
	Name        string `form:"name" json:"name" fieldName:"角色名" binding:"required,min=3,max=30"`
	Description string `form:"description" json:"description" fieldName:"描述" binding:"required,min=2,max=100"`
}

// ItemRole 根据 ID 查询角色表单
type ItemRole struct {
	ID int64 `form:"id" json:"id" fieldName:"ID" binding:"required"`
}
