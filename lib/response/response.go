package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response 返回封装
func (g *Gin) Response(httpCode, responseCode int, data ...interface{}) {
	if len(data) > 0 {
		g.C.JSON(httpCode, Response{
			Code: responseCode,
			Msg:  Msg(responseCode),
			Data: data[0],
		})
	} else {
		g.C.JSON(httpCode, Response{
			Code: responseCode,
			Msg:  Msg(responseCode),
			Data: "No data returned",
		})
	}
}

// SuccessResponse 成功返回
func (g *Gin) SuccessResponse(responseCode int, data ...interface{}) {
	g.Response(
		http.StatusOK,
		responseCode,
		data...,
	)
}

// ErrorResponse 失败返回
func (g *Gin) ErrorResponse(responseCode int, data ...interface{}) {
	g.Response(
		http.StatusOK,
		responseCode,
		data...,
	)
}

// ErrorMsgResponse 自定义错误信息返回
func (g *Gin) ErrorMsgResponse(responseMsg string, data ...interface{}) {
	if len(data) > 0 {
		g.C.JSON(http.StatusOK, Response{
			Code: CodeError,
			Msg:  responseMsg,
			Data: data[0],
		})
	} else {
		g.C.JSON(http.StatusOK, Response{
			Code: CodeError,
			Msg:  responseMsg,
			Data: "No data returned",
		})
	}
}

// SuccessMsgResponse 自定义成功信息返回
func (g *Gin) SuccessMsgResponse(responseMsg string, data ...interface{}) {
	if len(data) > 0 {
		g.C.JSON(http.StatusOK, Response{
			Code: CodeSuccess,
			Msg:  responseMsg,
			Data: data[0],
		})
	} else {
		g.C.JSON(http.StatusOK, Response{
			Code: CodeSuccess,
			Msg:  responseMsg,
			Data: "No data returned",
		})
	}
}
