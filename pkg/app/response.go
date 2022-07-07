package app

import (
	"net/http"

	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/gin-gonic/gin"
)

type Response struct {
	c *gin.Context
}

// State 状态码
type State struct {
	Code int         `json:"status_code"`    // 状态码，0-成功，其他值-失败
	Msg  string      `json:"status_msg"`     // 返回状态描述
	Data interface{} `json:"data,omitempty"` // 失败时返回空
}

type List struct {
	List interface{} `json:"list"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{c: ctx}
}

func (r *Response) Reply(err errcode.Err, datas ...interface{}) {
	var data interface{}
	if len(datas) > 0 {
		data = datas[0]
	}
	if err == nil {
		err = errcode.StatusOK
	} else {
		data = nil
	}
	r.c.JSON(http.StatusOK, State{
		Code: err.ECode(),
		Msg:  err.Error(),
		Data: data,
	})
}

func (r *Response) ReplyList(err errcode.Err, datas ...interface{}) {
	var data interface{}
	if len(datas) > 0 {
		data = datas[0]
	}
	if err == nil {
		err = errcode.StatusOK
	} else {
		data = nil
	}
	r.c.JSON(http.StatusOK, State{
		Code: err.ECode(),
		Msg:  err.Error(),
		Data: List{List: data},
	})
}
