package system

type AddMenu struct {
	ParentId  int64  `form:"parent_id" json:"parent_id" fieldName:"父级菜单 ID" binding:""`
	Name      string `form:"name" json:"name" fieldName:"菜单名称" binding:"required"`
	Title     string `form:"title" json:"title" fieldName:"菜单标题" binding:"required"`
	Path      string `form:"path" json:"path" fieldName:"路径" binding:"required"`
	Redirect  string `form:"redirect" json:"redirect" fieldName:"重定向路径" binding:""`
	Component string `form:"component" json:"component" fieldName:"组件" binding:"required"`
	Icon      string `form:"icon" json:"icon" fieldName:"图标" binding:""`
	Sort      int    `form:"sort" json:"sort" fieldName:"排序" binding:"required"`
	RoleList  string `form:"role_list" json:"role_list" fieldName:"角色 ID 组" binding:""`
	IsHidden  int    `form:"is_hidden" json:"is_hidden" fieldName:"显示状态" binding:""`
	IsAuth    int    `form:"is_auth" json:"is_auth" fieldName:"启用认证" binding:""`
}

type DelMenu struct {
	ID int64 `form:"id" json:"id" fieldName:"ID" binding:"required"`
}

type UpdMenu struct {
	ID        int64  `form:"id" json:"id" fieldName:"ID" binding:"required"`
	ParentId  int64  `form:"parent_id" json:"parent_id" fieldName:"父级菜单 ID" binding:""`
	Name      string `form:"name" json:"name" fieldName:"菜单名称" binding:"required"`
	Title     string `form:"title" json:"title" fieldName:"菜单标题" binding:"required"`
	Path      string `form:"path" json:"path" fieldName:"路径" binding:"required"`
	Redirect  string `form:"redirect" json:"redirect" fieldName:"重定向路径" binding:""`
	Component string `form:"component" json:"component" fieldName:"组件" binding:"required"`
	Icon      string `form:"icon" json:"icon" fieldName:"图标" binding:""`
	Sort      int    `form:"sort" json:"sort" fieldName:"排序" binding:"required"`
	RoleList  string `form:"role_list" json:"role_list" fieldName:"角色 ID 组" binding:""`
	IsHidden  int    `form:"is_hidden" json:"is_hidden" fieldName:"显示状态" binding:""`
	IsAuth    int    `form:"is_auth" json:"is_auth" fieldName:"启用认证" binding:""`
}

type ItemMenu struct {
	ID int64 `form:"id" json:"id" fieldName:"ID" binding:"required"`
}

type PatchIsAuthMenu struct {
	ID     int64 `form:"id" json:"id" fieldName:"ID" binding:"required"`
	IsAuth int   `form:"is_auth" json:"is_auth" fieldName:"启用认证" binding:""`
}

type PatchPathMenu struct {
	ID   int64  `form:"id" json:"id" fieldName:"ID" binding:"required"`
	Path string `form:"path" json:"path" fieldName:"路径" binding:"required"`
}

type PatchRedirectMenu struct {
	ID       int64  `form:"id" json:"id" fieldName:"ID" binding:"required"`
	Redirect string `form:"redirect" json:"redirect" fieldName:"重定向路径" binding:""`
}

type PatchIconMenu struct {
	ID   int64  `form:"id" json:"id" fieldName:"ID" binding:"required"`
	Icon string `form:"icon" json:"icon" fieldName:"图标" binding:""`
}

type MoveAuthMenu struct {
	ID       int64  `form:"id" json:"id" fieldName:"被拖拽菜单" binding:"required"`
	EndId    int64  `form:"end_id" json:"end_id" fieldName:"进入菜单" binding:"required"`
	Position string `form:"position" json:"position" fieldName:"位置" binding:"required"`
}
