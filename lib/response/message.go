package response

import (
	"fmt"
)

const (
	MsgSuccess = "成功"
	MsgError   = "失败"
)

// MsgList 消息 map
var MsgList = map[int]string{
	CodeSuccess:                    "成功",
	CodeError:                      "失败",
	CodeErrorInvalidParams:         "参数验证失败",
	CodeErrorAuthCheckTokenNull:    "请携带令牌 token",
	CodeErrorAuthCheckTokenFail:    "令牌认证失败",
	CodeErrorAuthCheckTokenTimeout: "令牌过期",
	CodeErrorAuthRoute:             "您无此操作权限",
	CodeErrorUserHasItem:           "用户名已存在",
	CodeErrorRoleHasItem:           "角色名已存在",
	CodeErrorRouteHasItem:          "此路由已存在",
	CodeErrorSystem:                "发生错误，请联系系统管理员",
	CodeErrorMongodb:               "数据库操作失败,请联系管理员",

	CodeSuccessAdd:  "添加成功",
	CodeSuccessDel:  "删除成功",
	CodeSuccessUpd:  "修改成功",
	CodeSuccessItem: "查询成功",
	CodeSuccessList: "列表查询成功",
	CodeErrorAdd:    "添加失败",
	CodeErrorDel:    "删除失败",
	CodeErrorUpd:    "修改失败",
	CodeErrorItem:   "查询失败",
	CodeErrorList:   "列表查询失败",

	CodeSuccessUserAdd:  MsgSuccessAddFormat(TypeUser),
	CodeSuccessUserDel:  MsgSuccessDelFormat(TypeUser),
	CodeSuccessUserUpd:  MsgSuccessUpdFormat(TypeUser),
	CodeSuccessUserItem: MsgSuccessItemFormat(TypeUser),
	CodeSuccessUserList: MsgSuccessListFormat(TypeUser),

	CodeSuccessRoleAdd:  MsgSuccessAddFormat(TypeRole),
	CodeSuccessRoleDel:  MsgSuccessDelFormat(TypeRole),
	CodeSuccessRoleUpd:  MsgSuccessUpdFormat(TypeRole),
	CodeSuccessRoleItem: MsgSuccessItemFormat(TypeRole),
	CodeSuccessRoleList: MsgSuccessListFormat(TypeRole),

	CodeErrorUserAdd:  MsgErrorAddFormat(TypeUser),
	CodeErrorUserDel:  MsgErrorDelFormat(TypeUser),
	CodeErrorUserUpd:  MsgErrorUpdFormat(TypeUser),
	CodeErrorUserItem: MsgErrorItemFormat(TypeUser),
	CodeErrorUserList: MsgErrorListFormat(TypeUser),

	CodeErrorRoleAdd:  MsgErrorAddFormat(TypeRole),
	CodeErrorRoleDel:  MsgErrorDelFormat(TypeRole),
	CodeErrorRoleUpd:  MsgErrorUpdFormat(TypeRole),
	CodeErrorRoleItem: MsgErrorItemFormat(TypeRole),
	CodeErrorRoleList: MsgErrorListFormat(TypeRole),

	CodeSuccessAuth: MsgSuccessActionFormat(ActionLogin),

	CodeErrorAuth: MsgErrorActionFormat(ActionLogin),
}

// Msg 获取消息
func Msg(code int) string {
	msg, ok := MsgList[code]
	if ok {
		return msg
	}
	return MsgError
}

// MsgSuccessFormat 成功消息格式化
func MsgSuccessFormat(typeName string, action string) string {
	return fmt.Sprintf("%s %s %s", typeName, action, MsgSuccess)
}

// MsgErrorFormat 失败消息格式化
func MsgErrorFormat(typeName string, action string) string {
	return fmt.Sprintf("%s %s %s", typeName, action, MsgError)
}

// MsgSuccessAddFormat 添加成功消息格式化
func MsgSuccessAddFormat(typeName string) string {
	return MsgSuccessFormat(typeName, ActionAdd)
}

// MsgSuccessDelFormat 删除成功消息格式化
func MsgSuccessDelFormat(typeName string) string {
	return MsgSuccessFormat(typeName, ActionDel)
}

// MsgSuccessUpdFormat 修改成功消息格式化
func MsgSuccessUpdFormat(typeName string) string {
	return MsgSuccessFormat(typeName, ActionUpd)
}

// MsgSuccessItemFormat 详情查询成功消息格式化
func MsgSuccessItemFormat(typeName string) string {
	return MsgSuccessFormat(typeName, ActionItem)
}

// MsgSuccessListFormat 列表查询添加成功消息格式化
func MsgSuccessListFormat(typeName string) string {
	return MsgSuccessFormat(typeName, ActionList)
}

// MsgErrorAddFormat 添加失败消息格式化
func MsgErrorAddFormat(typeName string) string {
	return MsgErrorFormat(typeName, ActionAdd)
}

// MsgErrorDelFormat 删除失败消息格式化
func MsgErrorDelFormat(typeName string) string {
	return MsgErrorFormat(typeName, ActionDel)
}

// MsgErrorUpdFormat 修改失败消息格式化
func MsgErrorUpdFormat(typeName string) string {
	return MsgErrorFormat(typeName, ActionUpd)
}

// MsgErrorItemFormat 详情查询失败消息格式化
func MsgErrorItemFormat(typeName string) string {
	return MsgErrorFormat(typeName, ActionItem)
}

// MsgErrorListFormat 列表查询添加失败消息格式化
func MsgErrorListFormat(typeName string) string {
	return MsgErrorFormat(typeName, ActionList)
}

// MsgSuccessActionFormat 操作成功消息格式化
func MsgSuccessActionFormat(action string) string {
	return fmt.Sprintf("%s %s", action, MsgSuccess)
}

// MsgErrorActionFormat 操作失败消息格式化
func MsgErrorActionFormat(action string) string {
	return fmt.Sprintf("%s %s", action, MsgError)
}
