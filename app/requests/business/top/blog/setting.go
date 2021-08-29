package blog

// UpdSetting 修改博客设置
type UpdSetting struct {
	Name  string `form:"name" json:"name" fieldName:"设置名" binding:"required,alphanum"`
	Value string `form:"value" json:"value" fieldName:"设置值" binding:"required"`
}

// NameSetting 读取博客设置
type NameSetting struct {
	Name string `form:"name" json:"name" fieldName:"设置名" binding:"required,alphanum"`
}
