package app

import (
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

//响应处理

type ResponseData struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Details []string    `json:"details"`
	Data    interface{} `json:"data"`
}

type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data interface{}) {
	r.Ctx.JSON(http.StatusOK, &ResponseData{
		Code: errcode.Success.Code(),
		Msg:  errcode.Success.Msg(),
		Data: data,
	})
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, &ResponseData{
		Code: errcode.Success.Code(),
		Msg:  errcode.Success.Msg(),
		Data: gin.H{
			"list": list,
			"pager": Pager{
				Page:      GetPage(r.Ctx),
				PageSize:  GetPageSize(r.Ctx),
				TotalRows: totalRows,
			},
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	r.Ctx.JSON(err.StatusCode(), &ResponseData{
		Code:    err.Code(),
		Msg:     err.Msg(),
		Details: err.Details(),
		Data:    nil,
	})
}
