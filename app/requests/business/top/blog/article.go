package blog

// AddArticle 添加文章
type AddArticle struct {
	Title      string `form:"title" json:"title" fieldName:"文章标题" binding:"required,min=3,max=30"`
	Content    string `form:"content" json:"content" fieldName:"文章内容" binding:"required,min=8"`
	ContentMd  string `form:"content_md" json:"content_md" fieldName:"文章内容(Markdown)" binding:""`
	CategoryId int64  `form:"category_id" json:"category_id" fieldName:"分类ID" binding:"required"`
	BannerId   int64  `form:"banner_id" json:"banner_id" fieldName:"文章横幅图ID" binding:""`
	EditorType int    `form:"editor_type" json:"editor_type" fieldName:"编辑器类型" binding:""`
}

// DelArticle 删除文章
type DelArticle struct {
	ID int64 `form:"id" json:"id" fieldName:"ID" binding:"required"`
}

// UpdArticle 修改文章
type UpdArticle struct {
	ID         int64  `form:"id" json:"id" fieldName:"ID" binding:"required"`
	Title      string `form:"title" json:"title" fieldName:"文章标题" binding:"required,min=3,max=20"`
	Content    string `form:"content" json:"content" fieldName:"文章内容" binding:"required,min=8"`
	ContentMd  string `form:"content_md" json:"content_md" fieldName:"文章内容(Markdown)" binding:""`
	CategoryId int64  `form:"category_id" json:"category_id" fieldName:"分类ID" binding:"required"`
	BannerId   int64  `form:"banner_id" json:"banner_id" fieldName:"文章横幅图ID" binding:""`
	EditorType int    `form:"editor_type" json:"editor_type" fieldName:"编辑器类型" binding:""`
}

// ItemArticle 查询文章
type ItemArticle struct {
	ID int64 `form:"id" json:"id" fieldName:"ID" binding:"required"`
}
