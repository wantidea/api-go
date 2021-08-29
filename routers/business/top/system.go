package top

import (
	SystemControllers "api-go/app/controllers/business/top/system"
	"api-go/app/middlewares"
	"github.com/gin-gonic/gin"
)

// AddSystemRoute 添加 系统 路由
func AddSystemRoute(r *gin.Engine) {
	systemR := r.Group(firstPrefix + "/system")

	adminC := &SystemControllers.Admin{}
	menuC := &SystemControllers.Menu{}
	roleC := &SystemControllers.Role{}
	routeC := &SystemControllers.Route{}
	dashboardC := &SystemControllers.Dashboard{}

	// 仪表台
	dashboardR := systemR.Group("/dashboard")
	{
		dashboardR.GET("list", dashboardC.List)
	}

	// 管理员
	adminR := systemR.Group("/admin")
	{
		adminR.POST("/login", adminC.Login)
	}
	// 需认证
	adminR.Use(middlewares.AdminAuthToken())
	{
		adminR.POST("/add", adminC.Add)
		adminR.DELETE("/del", adminC.Del)
		adminR.PUT("/upd", adminC.Upd)
		adminR.GET("/item", adminC.Item)
		adminR.GET("/list", adminC.List)
		adminR.GET("/info", adminC.Info)
		adminR.PATCH("/rePwd", adminC.RePwd)
	}

	// 菜单
	menuR := systemR.Group("/menu")
	{
		menuR.GET("/async", menuC.Async)
	}
	// 需认证
	menuR.Use(middlewares.AdminAuthToken())
	{
		menuR.POST("/add", menuC.Add)
		menuR.DELETE("/del", menuC.Del)
		menuR.PUT("/upd", menuC.Upd)
		menuR.GET("/item", menuC.Item)
		menuR.GET("/list", menuC.List)
		menuR.GET("/option", menuC.Option)
		menuR.GET("/treeList", menuC.TreeList)
		menuR.PATCH("/isAuth", menuC.PatchIsAuth)
		menuR.PATCH("/path", menuC.PatchPath)
		menuR.PATCH("/redirect", menuC.PatchRedirect)
		menuR.PATCH("/icon", menuC.PatchIcon)
		menuR.PATCH("/moveMenu", menuC.Move)
	}

	// 角色
	roleR := systemR.Group("/role")
	// 需认证
	roleR.Use(middlewares.AdminAuthToken())
	{
		roleR.POST("/add", roleC.Add)
		roleR.DELETE("/del", roleC.Del)
		roleR.PUT("/upd", roleC.Upd)
		roleR.GET("/item", roleC.Item)
		roleR.GET("/list", roleC.List)
		roleR.GET("/option", roleC.Option)
	}

	// 路由
	route := systemR.Group("/route")
	// 需认证
	route.Use(middlewares.AdminAuthToken())
	{
		route.POST("/add", routeC.Add)
		route.DELETE("/del", routeC.Del)
		route.PUT("/upd", routeC.Upd)
		route.GET("/item", routeC.Item)
		route.GET("/list", routeC.List)
	}
}
