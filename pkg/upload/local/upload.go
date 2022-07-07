package local

import (
	"mime/multipart"
	"os"
)

func SaveFile(fileType FileType, fileHeader *multipart.FileHeader) (string, error) {
	fileName, ext := GetFileName(fileHeader.Filename) // 获取加密文件名
	fileTypeor, err := checkContainExt(fileType, ext)
	if err != nil { // 判断文件类型是否合法
		return "", err
	}
	if !checkMaxSize(fileTypeor, fileHeader) { // 检查文件大小
		return "", FileSizeErr
	}
	uploadSavePath := fileTypeor.GetPath()
	if CheckSavePath(uploadSavePath) { // 检查保存路径是否存在
		if err := createSavePath(uploadSavePath, os.ModePerm); err != nil { // 创建保存路径
			return "", CreatePathErr
		}
	}
	if checkPermission(uploadSavePath) { // 检查权限
		return "", CompetenceErr
	}
	dst := uploadSavePath + "/" + fileName + ext // 加密文件名
	if err := saveFile(fileHeader, dst); err != nil {
		return "", err
	}
	accessUrl := fileTypeor.GetUrlPrefix() + "/" + fileName + ext
	return accessUrl, nil
}
