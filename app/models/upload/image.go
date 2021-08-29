package upload

import "gorm.io/plugin/soft_delete"

type Image struct {
	ID            int64                 `json:"id" gorm:"primarykey"`
	Url           string                `json:"url"`
	Path          string                `json:"path"`
	CreatedUserId int64                 `json:"created_user_id"`
	CreatedAt     int                   `json:"created_at"`
	UpdatedAt     int                   `json:"updated_at"`
	DeletedAt     soft_delete.DeletedAt `json:"deleted_at"`
}

func (m *Image) TableName() string {
	return "upload_image"
}
