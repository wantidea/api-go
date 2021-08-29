package system

import (
	"api-go/app/models/system"
	"api-go/app/services/redis"
	"api-go/lib/orm"
)

// AdminTotal 返回管理员数量
func AdminTotal() int64 {
	admin := &system.Admin{}
	var total int64
	orm.DB().Model(admin).Count(&total)
	return total
}

// AdminTotalOnRedis 返回管理员数量 读取 Redis
func AdminTotalOnRedis() int64 {
	total, err := redis.AdminTotal()
	if err != nil {
		total = AdminTotal()
		_ = redis.SetAdminTotal(total)
	}
	return total
}

// RoleTotal 返回管理员数量
func RoleTotal() int64 {
	role := &system.Role{}
	var total int64
	orm.DB().Model(role).Count(&total)
	return total
}

// RoleTotalOnRedis 返回管理员数量 读取 Redis
func RoleTotalOnRedis() int64 {
	total, err := redis.RoleTotal()
	if err != nil {
		total = RoleTotal()
		_ = redis.SetRoleTotal(total)
	}
	return total
}
