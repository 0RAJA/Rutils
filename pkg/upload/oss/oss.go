package upload

import (
	"mime/multipart"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// OSS 对象存储接口
type OSS interface {
	UploadFile(file *multipart.FileHeader, options []oss.Option) (string, string, error)
	DeleteFile(key ...string) (oss.DeleteObjectsResult, error)
}
