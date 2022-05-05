package app

import (
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

//响应处理

type Response struct {
	Ctx *gin.Context
}

type ResponseList struct {
	List  interface{} `json:"list"`
	Pager Pager       `json:"page"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data interface{}) {
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, ResponseList{
		List: list,
		Pager: Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

func (r *Response) ToErrorResponse(err errcode.Err) {
	r.Ctx.JSON(err.HCode(), err.Error())
}
