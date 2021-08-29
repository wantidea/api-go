package system

import (
	"gorm.io/plugin/soft_delete"
)

type Menu struct {
	ID            int64                 `json:"id" gorm:"primarykey"`
	ParentId      int64                 `json:"parent_id"`
	Level         int                   `json:"level"`
	Name          string                `json:"name"`
	Title         string                `json:"title"`
	Path          string                `json:"path"`
	Redirect      string                `json:"redirect"`
	Component     string                `json:"component"`
	Icon          string                `json:"icon"`
	Sort          int                   `json:"sort"`
	RoleList      string                `json:"role_list"`
	IsHidden      int                   `json:"is_hidden"`
	IsAuth        int                   `json:"is_auth"`
	CreatedUserId int64                 `json:"created_user_id"`
	CreatedAt     int                   `json:"created_at"`
	UpdatedAt     int                   `json:"updated_at"`
	DeletedAt     soft_delete.DeletedAt `json:"deleted_at"`
}

func (m *Menu) TableName() string {
	return "system_menu"
}
