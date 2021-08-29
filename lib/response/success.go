package response

// 成功编码
const (
	// 错误码统一格式 A-BB-CCC
	// A: 1-2-3-成功
	// A: 4-客户端异常 请求方式错误、请求参数错误、没有权限、没有登录、上传文件过大、重复提交抢锁失败、并发过大触发限流等
	// A: 5-服务端异常 依赖的第三方业务系统异常：比如调用第三方系统超时；第三方系统抛出异常；第三方业务系统限流等
	// A: 6-基础中间件异常 如 MySql、Redis、Mongodb、MQ 等基础中间件出现连接超时、连接池满、访问失败等
	// A: 7-数据问题 数据不一致、记录不存在、主键冲突、字段不能为空等等
	// BB: 模块名称
	// CCC: 具体错误编号 自增即可

	// 通用成功 00
	CodeSuccess     = 100000
	CodeSuccessAdd  = 100001
	CodeSuccessDel  = 100002
	CodeSuccessUpd  = 100003
	CodeSuccessItem = 100004
	CodeSuccessList = 100005
	CodeSuccessAuth = 100006

	// 用户 01
	CodeSuccessUserAdd  = 101001
	CodeSuccessUserDel  = 101002
	CodeSuccessUserUpd  = 101003
	CodeSuccessUserItem = 101004
	CodeSuccessUserList = 101005

	// 角色 02
	CodeSuccessRoleAdd  = 102001
	CodeSuccessRoleDel  = 102002
	CodeSuccessRoleUpd  = 102003
	CodeSuccessRoleItem = 102004
	CodeSuccessRoleList = 102005

	// 路由 03
	CodeSuccessRouteAdd  = 103001
	CodeSuccessRouteDel  = 103002
	CodeSuccessRouteUpd  = 103003
	CodeSuccessRouteItem = 103004
	CodeSuccessRouteList = 103005

	// 文件上传 04
	CodeSuccessUploadImage     = 104001
	CodeSuccessUploadImageList = 104002
)
