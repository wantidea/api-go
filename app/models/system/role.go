package system

import (
	"api-go/lib/orm"
	"errors"
	"gorm.io/plugin/soft_delete"
)

type Role struct {
	ID            int64                 `json:"id" gorm:"primarykey"`
	Name          string                `json:"name"`
	Description   string                `json:"description"`
	CreatedUserId int64                 `json:"created_user_id"`
	CreatedAt     int                   `json:"created_at"`
	UpdatedAt     int                   `json:"updated_at"`
	DeletedAt     soft_delete.DeletedAt `json:"deleted_at"`
}

func (m *Role) TableName() string {
	return "system_role"
}

// CreateCheckName 新增 检查名称
func (m *Role) CreateCheckName() error {
	var count int64
	orm.DB().Model(Role{}).Where("name = ?", m.Name).Count(&count)

	if count > 0 {
		return errors.New("该名称已存在")
	} else {
		return nil
	}
}

// UpdateCheckName 修改 检查名称
func (m *Role) UpdateCheckName() error {
	var count int64
	orm.DB().Model(Role{}).Where("id != ? AND name = ?", m.ID, m.Name).Count(&count)

	if count > 0 {
		return errors.New("该名称已存在")
	} else {
		return nil
	}
}
