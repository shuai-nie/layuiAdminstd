package service

import (
	"layuiAdminstd/global"
	"layuiAdminstd/pkg/upload"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	uploadSavePath := upload.GetSavePath()
	dst := uploadSavePath + "/" + fileName
	if !upload.CheckContaionExt(fileType, fileName) {
		//
	}

	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			//
		}
	}

	if upload.CheckMaxSize(fileType, file) {
		//
	}

	if upload.CheckPermission(uploadSavePath) {
		//
	}

	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl:  accessUrl}, nil
}
