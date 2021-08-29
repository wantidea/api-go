package article

import "gorm.io/plugin/soft_delete"

type Category struct {
	ID            int64                 `json:"id" gorm:"primarykey"`
	Name          string                `json:"name"`
	CreatedUserId int64                 `json:"created_user_id"`
	CreatedAt     int                   `json:"created_at"`
	UpdatedAt     int                   `json:"updated_at"`
	DeletedAt     soft_delete.DeletedAt `json:"deleted_at"`
}

func (m *Category) TableName() string {
	return "blog_article_category"
}
