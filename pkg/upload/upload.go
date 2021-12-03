package upload

import (
	"mime/multipart"
	"os"
)

func SaveFile(fileType FileType, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	fileName := GetFileName(fileHeader.Filename) //获取文件名
	if !checkContainExt(fileType, fileName) {    //判断文件类型是否合法
		return "", ExtErr
	}
	if !checkMaxSize(fileType, file) { //检查文件大小
		return "", FileSizeErr
	}
	uploadSavePath := GetSavePath()
	if CheckSavePath(uploadSavePath) { //检查保存路径是否存在
		if err := createSavePath(uploadSavePath, os.ModePerm); err != nil { //创建保存路径
			return "", CreatePathErr
		}
	}
	if checkPermission(uploadSavePath) { //检查权限
		return "", CompetenceErr
	}
	dst := uploadSavePath + "/" + fileName //加密文件名
	if err := saveFile(fileHeader, dst); err != nil {
		return "", err
	}
	accessUrl := ServerInit.ServerUrl + "/" + fileName
	return accessUrl, nil
}
