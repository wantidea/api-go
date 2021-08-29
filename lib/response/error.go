package response

// 错误编码
const (
	// 错误码统一格式 A-BB-CCC
	// A: 1-2-3-成功
	// A: 4-客户端异常 请求方式错误、请求参数错误、没有权限、没有登录、上传文件过大、重复提交抢锁失败、并发过大触发限流等
	// A: 5-服务端异常 依赖的第三方业务系统异常：比如调用第三方系统超时；第三方系统抛出异常；第三方业务系统限流等
	// A: 6-基础中间件异常 如 MySql、Redis、Mongodb、MQ 等基础中间件出现连接超时、连接池满、访问失败等
	// A: 7-数据问题 数据不一致、记录不存在、主键冲突、字段不能为空等等
	// BB: 模块名称
	// CCC: 具体错误编号 自增即可

	// 通用失败 00
	CodeError                      = 400000
	CodeErrorAdd                   = 400001
	CodeErrorDel                   = 400002
	CodeErrorUpd                   = 400003
	CodeErrorItem                  = 400004
	CodeErrorList                  = 400005
	CodeErrorInvalidParams         = 400006
	CodeErrorHasItem               = 400007
	CodeErrorAuth                  = 400008
	CodeErrorAuthCheckTokenNull    = 400009
	CodeErrorAuthCheckTokenFail    = 400010
	CodeErrorAuthCheckTokenTimeout = 400011
	CodeErrorAuthRoute             = 400012
	CodeErrorSystem                = 400013

	// 用户 01
	CodeErrorUserAdd     = 401001
	CodeErrorUserDel     = 401002
	CodeErrorUserUpd     = 401003
	CodeErrorUserItem    = 401004
	CodeErrorUserList    = 401005
	CodeErrorUserHasItem = 401006

	// 角色 02
	CodeErrorRoleAdd      = 402001
	CodeErrorRoleDel      = 402002
	CodeErrorRoleUpd      = 402003
	CodeErrorRoleItem     = 402004
	CodeErrorRoleList     = 402005
	CodeErrorRoleHasItem  = 402006
	CodeErrorRouteHasItem = 402007

	// 路由 03
	CodeErrorRouteAdd  = 403001
	CodeErrorRouteDel  = 403002
	CodeErrorRouteUpd  = 403003
	CodeErrorRouteItem = 403004
	CodeErrorRouteList = 403005

	// 文件上传 04
	CodeErrorUploadImage     = 404001
	CodeErrorUploadImageList = 404002

	// Mongodb 01
	CodeErrorMongodb = 601000
)
