package blog

// ItemArticle 查询文章
type ItemArticle struct {
	ID int64 `form:"id" json:"id" fieldName:"ID" binding:"required"`
}
