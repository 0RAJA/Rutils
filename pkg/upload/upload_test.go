package upload_test

import (
	"github.com/0RAJA/Rutils/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestSaveFile(t *testing.T) {
	image := upload.NewFile("image", []string{".PNG"}, 1024*1024*20, "http://127.0.0.1:2333/static/images", "./images")
	//按上面那样子添加支持的文件类型和其对应的文件后缀
	upload.Init(image)
	routine := gin.Default()
	routine.Static("/static/", "/home/raja/workspace/go/src/Rutils/pkg/upload")
	routine.POST("/upload", func(c *gin.Context) {
		_, fileHeader, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		fileType := upload.FileType(c.PostForm("filetype"))
		url, err := upload.SaveFile(fileType, fileHeader)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"url": url})
	})
	if err := routine.Run(":2333"); err != nil {
		panic(err)
	}
}
