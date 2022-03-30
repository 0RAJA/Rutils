package app

import (
	"github.com/gin-gonic/gin"
)

//分页处理

var (
	DefaultPageSize int32
	MaxPageSize     int32
	PageKey         string //url中page关键字
	PageSizeKey     string //pagesize关键字
)

// Init 初始化默认页数大小和最大页数限制以及查询的关键字
func Init(defaultPageSize, maxPageSize int32, pageKey, pageSizeKey string) {
	DefaultPageSize = defaultPageSize
	MaxPageSize = maxPageSize
	PageKey = pageKey
	PageSizeKey = pageSizeKey
}

type Pager struct {
	Page      int32 `json:"page,omitempty"`
	PageSize  int32 `json:"page_size,omitempty"`
	TotalRows int   `json:"total_rows"`
}

// GetPage 获取页数
func GetPage(c *gin.Context) int32 {
	page := StrTo(c.Query(PageKey)).MustInt32()
	if page <= 0 {
		return 1
	}
	return page
}

// GetPageSize 获取pageSize
func GetPageSize(c *gin.Context) int32 {
	pageSize := StrTo(c.Query(PageSizeKey)).MustInt32()
	if pageSize <= 0 {
		return DefaultPageSize
	}
	if pageSize > MaxPageSize {
		return MaxPageSize
	}
	return pageSize
}

// GetPageOffset 通过page和pageSize获取偏移值
func GetPageOffset(page, pageSize int32) (result int32) {
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return
}
