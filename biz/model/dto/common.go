package dto

import (
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"webchat_be/biz/model/errs"
)

type CommonResp struct {
	Success bool        `json:"success"`
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResp(c *app.RequestContext, data interface{}) {
	c.JSON(http.StatusOK, &CommonResp{
		Success: true,
		Code:    errs.Success.Code(),
		Message: errs.Success.Msg(),
		Data:    data,
	})
}

func FailResp(c *app.RequestContext, bizErr errs.Error) {
	c.JSON(http.StatusOK, &CommonResp{
		Success: false,
		Code:    bizErr.Code(),
		Message: bizErr.Msg(),
	})
}

func AbortWithErr(c *app.RequestContext, bizErr errs.Error, code int) {
	c.AbortWithStatusJSON(code, &CommonResp{
		Success: false,
		Code:    bizErr.Code(),
		Message: bizErr.Msg(),
	})
}
