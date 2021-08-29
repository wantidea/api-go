package redis

import (
	"api-go/lib/redis"
	"context"
	"strconv"
	"strings"
	"time"
)

const (
	// 数字
	IAdminTotal   = "SYSTEM:ADMIN:TOTAL"   // 管理员数量
	IRoleTotal    = "SYSTEM:ROLE:TOTAL"    // 角色数量
	IArticleTotal = "SYSTEM:ARTICLE:TOTAL" // 文章数量

	// 字符串
	SMenuTree     = "SYSTEM:MENU:TREE"       // 树形菜单列表
	SMenuTreeRole = "SYSTEM:MENU:ROLE:role_" // 树形菜单验证列表 角色

	// 过期时间设置 秒 0-无过期时间
	IAdminTotalExpirationTime   = 3600 * time.Second // 管理员数量
	IRoleTotalExpirationTime    = 3600 * time.Second // 角色数量
	IArticleTotalExpirationTime = 3600 * time.Second // 文章数量
	SMenuTreeExpirationTime     = 0                  // 树形菜单列表
	SMenuTreeRoleExpirationTime = 0                  // 树形菜单验证列表 角色
)

// AdminTotal 读取管理员数量
func AdminTotal() (int64, error) {
	return redis.Rdb().Get(context.Background(), IAdminTotal).Int64()
}

// SetAdminTotal 设置管理员数量
func SetAdminTotal(total int64) error {
	return redis.Rdb().Set(context.Background(), IAdminTotal, total, IAdminTotalExpirationTime).Err()
}

// RoleTotal 读取角色数量
func RoleTotal() (int64, error) {
	return redis.Rdb().Get(context.Background(), IRoleTotal).Int64()
}

// SetRoleTotal 设置角色数量
func SetRoleTotal(total int64) error {
	return redis.Rdb().Set(context.Background(), IRoleTotal, total, IRoleTotalExpirationTime).Err()
}

// ArticleTotal 读取文章数量
func ArticleTotal() (int64, error) {
	return redis.Rdb().Get(context.Background(), IArticleTotal).Int64()
}

// SetArticleTotal 设置文章数量
func SetArticleTotal(total int64) error {
	return redis.Rdb().Set(context.Background(), IArticleTotal, total, IArticleTotalExpirationTime).Err()
}

// MenuTree 读取树形菜单列表
func MenuTree() ([]byte, error) {
	return redis.Rdb().Get(context.Background(), SMenuTree).Bytes()
}

// SetMenuTree 设置树形菜单列表
func SetMenuTree(treeMenu []byte) error {
	return redis.Rdb().Set(context.Background(), SMenuTree, treeMenu, SMenuTreeExpirationTime).Err()
}

// MenuTreeOnAuthByRoleId 设置树形菜单验证列表 角色
func MenuTreeOnAuthByRoleId(roleId int64) ([]byte, error) {
	var build strings.Builder
	build.WriteString(SMenuTreeRole)
	build.WriteString(strconv.FormatInt(roleId, 10))
	key := build.String()
	return redis.Rdb().Get(context.Background(), key).Bytes()
}

// SetMenuTreeOnAuthByRoleId 设置树形菜单验证列表 角色
func SetMenuTreeOnAuthByRoleId(treeMenu []byte, roleId int64) error {
	var build strings.Builder
	build.WriteString(SMenuTreeRole)
	build.WriteString(strconv.FormatInt(roleId, 10))
	key := build.String()
	return redis.Rdb().Set(context.Background(), key, treeMenu, SMenuTreeRoleExpirationTime).Err()
}
