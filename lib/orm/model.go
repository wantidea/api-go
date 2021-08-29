package orm

import (
	"gorm.io/plugin/soft_delete"
)

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt int
	UpdatedAt int
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name"`
}
