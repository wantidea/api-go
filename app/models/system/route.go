package system

import (
	"gorm.io/plugin/soft_delete"
)

type Route struct {
	ID            int64                 `json:"id" gorm:"primarykey"`
	Name          string                `json:"name"`
	Uri           string                `json:"uri"`
	Method        string                `json:"method"`
	RoleList      string                `json:"role_list"`
	IsAuth        int8                  `json:"is_auth"`
	CreatedUserId int64                 `json:"created_user_id"`
	CreatedAt     int                   `json:"created_at"`
	UpdatedAt     int                   `json:"updated_at"`
	DeletedAt     soft_delete.DeletedAt `json:"deleted_at"`
}

func (m *Route) TableName() string {
	return "system_route"
}
