package article

// AddCategory 添加分类
type AddCategory struct {
	Name string `form:"name" json:"name" fieldName:"分类名称" binding:"required"`
}

// DelCategory 删除分类
type DelCategory struct {
	ID int64 `form:"id" json:"id" fieldName:"ID" binding:"required"`
}

// UpdCategory 修改分类
type UpdCategory struct {
	ID   int64  `form:"id" json:"id" fieldName:"ID" binding:"required"`
	Name string `form:"name" json:"name" fieldName:"分类名称" binding:"required"`
}

// ItemCategory 查询分类
type ItemCategory struct {
	ID int64 `form:"id" json:"id" fieldName:"ID" binding:"required"`
}

// ListArticleCategory 查询分类
type ListArticleCategory struct {
	ID int64 `form:"id" json:"id" fieldName:"ID" binding:"required"`
}
