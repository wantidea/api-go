package system

import (
	"api-go/lib/md5"
	"api-go/lib/orm"
	"errors"
	"gorm.io/plugin/soft_delete"
)

type Admin struct {
	ID            int64                 `json:"id" gorm:"primarykey"`
	Name          string                `json:"name"`
	Password      string                `json:"password"`
	AvatarId      int64                 `json:"avatar_id"`
	RoleId        int64                 `json:"role_id"`
	CreatedUserId int64                 `json:"created_user_id"`
	CreatedAt     int                   `json:"created_at"`
	UpdatedAt     int                   `json:"updated_at"`
	DeletedAt     soft_delete.DeletedAt `json:"deleted_at"`
}

func (m *Admin) TableName() string {
	return "system_admin"
}

// CreateCheckName 新增 检查名称
func (m *Admin) CreateCheckName() error {
	var count int64
	orm.DB().Model(Admin{}).Where("name = ?", m.Name).Count(&count)

	if count > 0 {
		return errors.New("该名称已存在")
	} else {
		return nil
	}
}

// UpdateCheckName 修改 检查名称
func (m *Admin) UpdateCheckName() error {
	var count int64
	orm.DB().Model(Admin{}).Where("id != ? AND name = ?", m.ID, m.Name).Count(&count)

	if count > 0 {
		return errors.New("该名称已存在")
	} else {
		return nil
	}
}

// Auth 验证用户名密码
func (m *Admin) Auth() bool {
	result := orm.DB().Where("name = ? AND password = ?", m.Name, md5.ToMD5Salt(m.Password)).First(m)
	if result.Error != nil || result.RowsAffected <= 0 {
		return false
	}
	return true
}

// CreateAdminCheckName 检查名称是否存在
func CreateAdminCheckName(name string) error {
	admin := &Admin{
		Name: name,
	}
	return admin.CreateCheckName()
}
