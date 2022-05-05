package upload

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type FileType string

type FileTypeor interface {
	GetType() FileType
	GetSuffix() []string
	GetMaxSize() int
	GetUrlPrefix() string
	GetPath() string
}

var UploadManager = make(uploadManager)

type uploadManager map[FileType]FileTypeor

func (u uploadManager) addFileType(fileType FileTypeor) {
	if _, ok := u[fileType.GetType()]; ok {
		panic(RepeatedFileTypeErr)
	}
	u[fileType.GetType()] = fileType
}

var (
	ExtErr              = errors.New("file suffix is not supported")
	FileSizeErr         = errors.New("exceeded maximum file limit")
	CreatePathErr       = errors.New("failed to create save directory")
	CompetenceErr       = errors.New("insufficient file permissions")
	RepeatedFileTypeErr = errors.New("DuplicateFileType")
)

func Init(typeors ...FileTypeor) {
	for _, typeor := range typeors {
		UploadManager.addFileType(typeor)
	}
}

// GetFileExt 获取后缀
func GetFileExt(name string) string {
	return path.Ext(name)
}

// GetFileName 加密文件名
func GetFileName(fileName string) (string, string) {
	ext := GetFileExt(fileName)
	m := md5.New()
	t := strconv.Itoa(int(time.Now().Unix()))
	m.Write([]byte(fileName + t))
	return hex.EncodeToString(m.Sum(nil)), ext
}

//检查文件

// CheckSavePath 检查保存目录是否存在
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return errors.Is(err, os.ErrNotExist) //文件不存在
}

// checkContainExt 检查文件后缀是否包含在约定的后缀配置项中
func checkContainExt(t FileType, ext string) (fileTypeor FileTypeor, err error) {
	ext = strings.ToUpper(ext)
	fileTypeor, ok := UploadManager[t]
	if !ok {
		return nil, ExtErr
	}
	for _, suffix := range fileTypeor.GetSuffix() {
		if suffix == ext {
			return fileTypeor, nil
		}
	}
	return nil, ExtErr
}

// checkMaxSize 检查文件大小是否超出最大大小限制
func checkMaxSize(t FileTypeor, f *multipart.FileHeader) bool {
	return t.GetMaxSize() >= int(f.Size)
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
