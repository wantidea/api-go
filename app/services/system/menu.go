package system

import (
	SystemModels "api-go/app/models/system"
	RabbitMQServices "api-go/app/services/rabbitmq"
	RedisServices "api-go/app/services/redis"
	"api-go/lib/orm"
	"encoding/json"
)

type MenuTree struct {
	ID        int64       `json:"id" gorm:"primarykey"`
	ParentId  int64       `json:"parent_id"`
	Level     int         `json:"level"`
	Name      string      `json:"name"`
	Title     string      `json:"title"`
	Path      string      `json:"path"`
	Redirect  string      `json:"redirect"`
	Component string      `json:"component"`
	Icon      string      `json:"icon"`
	Sort      int         `json:"sort"`
	RoleList  string      `json:"role_list"`
	IsHidden  int         `json:"is_hidden"`
	IsAuth    int         `json:"is_auth"`
	Children  []*MenuTree `json:"children" gorm:"-"`
}

// BuildTreeMenu 打包树形菜单列表
func BuildTreeMenu(menuList []*MenuTree, parentId int64) []*MenuTree {
	var tmp []*MenuTree

	for _, menuItem := range menuList {
		if menuItem.ParentId == parentId {
			menuItem.Children = BuildTreeMenu(menuList, menuItem.ID)
			tmp = append(tmp, menuItem)
		}
	}

	return tmp
}

// TreeMenu 返回树形菜单列表
func TreeMenu() []*MenuTree {
	var result []*MenuTree
	orm.DB().Model(&SystemModels.Menu{}).
		Select(
			"id",
			"parent_id",
			"level",
			"name",
			"title",
			"path",
			"redirect",
			"component",
			"icon",
			"sort",
			"role_list",
			"is_hidden",
			"is_auth",
		).
		Order("sort asc").
		Find(&result)

	return BuildTreeMenu(result, 0)
}

// TreeMenuOnAuthByRoleId 返回经过验证的角色树形菜单列表
func TreeMenuOnAuthByRoleId(roleId int64) []*MenuTree {
	var result []*MenuTree
	orm.DB().Model(&SystemModels.Menu{}).
		Select(
			"id",
			"parent_id",
			"level",
			"name",
			"title",
			"path",
			"redirect",
			"component",
			"icon",
			"sort",
			"role_list",
			"is_hidden",
			"is_auth",
		).
		Where("(is_auth = 1 and find_in_set(?,role_list)) OR (is_auth = 0)", roleId).
		Order("sort asc").
		Find(&result)

	return BuildTreeMenu(result, 0)
}

// TreeMenuOnRedis 获取经过验证的角色树形菜单列表
func TreeMenuOnRedis() ([]*MenuTree, error) {
	var result []*MenuTree
	resultRedis, err := RedisServices.MenuTree()
	if err != nil {
		result = TreeMenu()
		resultJson, _ := json.Marshal(result)
		_ = RedisServices.SetMenuTree(resultJson)
		return result, nil
	}

	_ = json.Unmarshal(resultRedis, &result)
	return result, nil
}

// RefreshTreeMenu 刷新树形菜单
func RefreshTreeMenu() error {
	result, _ := json.Marshal(TreeMenu())
	return RedisServices.SetMenuTree(result)
}

// TreeMenuOnAuthRedisByRoleId 获取经过验证的角色树形菜单列表
func TreeMenuOnAuthRedisByRoleId(roleId int64) ([]*MenuTree, error) {
	var result []*MenuTree
	resultRedis, err := RedisServices.MenuTreeOnAuthByRoleId(roleId)
	if err != nil {
		menuTree := TreeMenuOnAuthByRoleId(roleId)
		menuTreeJson, _ := json.Marshal(menuTree)
		_ = RedisServices.SetMenuTreeOnAuthByRoleId(menuTreeJson, roleId)
		return menuTree, nil
	}

	_ = json.Unmarshal(resultRedis, &result)
	return result, nil
}

// RefreshTreeMenuOnAuthByRoleId 刷新经过验证的角色树形菜单列表
func RefreshTreeMenuOnAuthByRoleId(roleId int64) error {
	treeMenu := TreeMenuOnAuthByRoleId(roleId)
	resultJson, _ := json.Marshal(treeMenu)
	return RedisServices.SetMenuTreeOnAuthByRoleId(resultJson, roleId)
}

// RefreshTreeMenuOnAll 刷新所有树形菜单
func RefreshTreeMenuOnAll() error {
	// 更新树形菜单
	_ = RefreshTreeMenu()

	// 更新所有角色动态菜单
	var roleList []*SystemModels.Role
	orm.DB().Model(&SystemModels.Role{}).
		Select("id").
		Find(&roleList)

	for i := 0; i < len(roleList); i++ {
		_ = RefreshTreeMenuOnAuthByRoleId(roleList[i].ID)
	}

	return nil
}

type TreeMenuConsume struct {
}

// AddRefreshTreeMenuProducer 添加更新消息队列
func AddRefreshTreeMenuProducer() error {
	return RabbitMQServices.RegisterProducer(
		RabbitMQServices.ERedis,
		RabbitMQServices.ExchangeDirect,
		RabbitMQServices.RRedisTreeMenu,
		RabbitMQServices.MRedisTreeMenuRefresh,
	)
}

// ListenRefreshTreeMenuPromise 监听消费
func ListenRefreshTreeMenuPromise() {
	t := &TreeMenuConsume{}
	RabbitMQServices.RegisterReceiver(
		RabbitMQServices.ERedis,
		RabbitMQServices.ExchangeDirect,
		RabbitMQServices.RRedisTreeMenu,
		t,
	)
}

// MsgConsume 消费内容
func (t *TreeMenuConsume) MsgConsume(dataByte []byte) error {
	return RefreshTreeMenuOnAll()
}
