package upload

import (
	"Rutils/pkg/util"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

var ImageInit *ImageStruct

var ServerInit *ServerStruct

type ImageStruct struct {
	ImageAllowExits  []string
	SaveImageMaxSize int
}

type ServerStruct struct {
	SavePath  string
	ServerUrl string
}

const TypeImage FileType = iota + 1

var (
	ExtErr        = errors.New("file suffix is not supported")
	FileSizeErr   = errors.New("exceeded maximum file limit")
	CreatePathErr = errors.New("failed to create save directory")
	CompetenceErr = errors.New("insufficient file permissions")
)

func Init(serverStruct *ServerStruct, imageStruct *ImageStruct) {
	ServerInit = serverStruct
	ImageInit = imageStruct
}

// GetFileExt 获取后缀
func GetFileExt(name string) string {
	return path.Ext(name)
}

// GetSavePath 获取默认保存路径
func GetSavePath() string {
	return ServerInit.SavePath
}

// GetFileName 加密文件名
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext) //去除文件后缀
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

//检查文件

// CheckSavePath 检查保存目录是否存在
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return errors.Is(err, os.ErrNotExist) //文件不存在
}

// checkContainExt 检查文件后缀是否包含在约定的后缀配置项中
func checkContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range ImageInit.ImageAllowExits {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}
	return false
}

// checkMaxSize 检查文件大小是否超出最大大小限制
func checkMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size < ImageInit.SaveImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

// checkPermission 检查文件权限是否足够
func checkPermission(dst string) bool {
	_, err := os.Stat(dst)
	return errors.Is(err, os.ErrPermission)
}

//文件写入/创建的相关操作

// createSavePath 创建保存路径 perm 表示目录权限
func createSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

func saveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	out, err := os.Create(dst) //创建文件
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src) //写入文件
	return err
}
