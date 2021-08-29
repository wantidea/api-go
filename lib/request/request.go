package request

import (
	"api-go/lib/response"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// CheckForm 表单验证
func CheckForm(c *gin.Context, appG *response.Gin, form interface{}) error {
	err := c.ShouldBind(form)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//appG.ErrorResponse(
			//	response.CodeErrorInvalidParams,
			//	fmt.Sprintf("Error: %s", err.Error()),
			//)
			appG.ErrorMsgResponse(fmt.Sprintf("Error: %s", err.Error()))
		} else {
			//appG.ErrorResponse(
			//	response.CodeErrorInvalidParams,
			//	RemoveTopStruct(errs.Translate(Trans)),
			//)
			appG.ErrorMsgResponse(fmt.Sprintf("参数验证错误：%s", FirstError(errs.Translate(Trans))))
		}
		return errors.New(err.Error())
	}
	return nil
}
