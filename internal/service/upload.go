package service

import (
	"github.com/gin-gonic/gin"
	"io"
	"layuiAdminstd/global"
	"layuiAdminstd/pkg/upload"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"errors"
)

type FileInfo struct {
	Name string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported.")
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit.")
	}

	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory.")
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions.")
	}

	dst := uploadSavePath + "/" + fileName
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}


func UploadResource(c *gin.Context) {
	saveDirParam := c.PostForm("saveDir")
	fileNameParam := c.PostForm("fileName")

	var saveDir = ""
	var saveName = ""
	var savePath = ""
	file, header, errFile := c.Request.FormFile("file")
	if errFile != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "请选择文件",
			"dir": saveDir,
			"name": saveName,
			"path": savePath,
		})
		return
	}

	// 回去上传文件后缀(类型）
	uploadFileNameWithSuffix := path.Base(header.Filename)
	uploadFileType := path.Ext(uploadFileNameWithSuffix)
	// 文件保存目录
	saveDir = "" + saveDirParam
	// 保存的文件名称
	saveName = fileNameParam + uploadFileType
	savePath = saveDir + "/" + saveName
	// 打开目录
	localFileInfo, fileStatErr := os.Stat(saveDir)
	// 目录不存在
	if fileStatErr != nil || !localFileInfo.IsDir() {
		// 创建目录
		errByMkdirAllDir := os.MkdirAll(saveDir, 0755)
		if errByMkdirAllDir != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"dir": saveDir,
				"name": saveName,
				"path":savePath,
				"message":"创建目录失败",
			})
			return
		}
	}
	// 上传文件前，先删除该资源之前上传过的资源文件
	// 编辑-重新选择文件，需要先删除资源之前上传过的资源文件
	// 该代码执行的条件——上传的名称是唯一的，否则会出现误删
	// 获取文件的前缀
	// fileNameOnly := fileNameParam
	// deleteFileWithName(fileNameOnly, saveDir)
	// deleteFileWithName(fileNameOnly, model.WebConfig.ResourcePath + "/" + model.WebConfig.WebConvertToPath)

	out, err := os.Create(savePath)
	if err != nil {

	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"dir": saveDir,
			"name": saveName,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"dir": saveDir,
		"name": saveName,
		"path": savePath,
		"message": "上传成功",
	})
}