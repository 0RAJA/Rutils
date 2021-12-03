package app

import (
	"Rutils/pkg/app/convert"
	"github.com/gin-gonic/gin"
)

//分页处理

var (
	DefaultPageSize int
	MaxPageSize     int
)

func Init(defaultPageSize, maxPageSize int) {
	DefaultPageSize = defaultPageSize
	MaxPageSize = maxPageSize
}

type Pager struct {
	Page      int `json:"page,omitempty"`
	PageSize  int `json:"page_size,omitempty"`
	TotalRows int `json:"total_rows,omitempty"`
}

func NewPager(page int, pageSize int) *Pager {
	return &Pager{Page: page, PageSize: pageSize}
}

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return DefaultPageSize
	}
	if pageSize > MaxPageSize {
		return MaxPageSize
	}
	return pageSize
}

func GetPageOffset(page, pageSize int) (result int) {
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return
}
