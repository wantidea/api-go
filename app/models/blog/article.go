package blog

import "gorm.io/plugin/soft_delete"

type Article struct {
	ID            int64                 `json:"id" gorm:"primarykey"`
	Title         string                `json:"title"`
	CategoryId    int64                 `json:"category_id"`
	BannerId      int64                 `json:"banner_id"`
	EditorType    int                   `json:"editor_type"`
	Content       string                `json:"content"`
	ContentMd     string                `json:"content_md"`
	LookTotal     int                   `json:"look_total"`
	CreatedUserId int64                 `json:"created_user_id"`
	CreatedAt     int                   `json:"created_at"`
	UpdatedAt     int                   `json:"updated_at"`
	DeletedAt     soft_delete.DeletedAt `json:"deleted_at"`
}

func (m *Article) TableName() string {
	return "blog_article"
}
