package qiniu

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

func UploadToQiNiu(file *multipart.FileHeader) (int, string) {
	accessKey := ""
	accessSerctKey := ""
	bucket := ""
	imgUrl := ""

	src, err := file.Open()
	if err != nil {
		return 100111, err.Error()
	}
	defer src.Close()

	putPlcy := storage.PutPolicy{
		Scope: bucket,
	}

	mac := qbox.NewMac(accessKey, accessSerctKey)

	// 获取上传凭证
	upToken := putPlcy.UploadToken(mac)

	// 配置参数
	cfg := storage.Config{
		Zone: &storage.ZoneHuadong,//华南区
		UseCdnDomains: false,
		UseHTTPS: false, // 非https
	}

	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PfopRet{} // 上传后返回的结果
	putExtra := storage.PutExtra{} // 额外参数

	key := "image/" + file.Filename // 上传路径，如果当前目录中已存在相同文件，则返回上传失败错误
	err = formUploader.Put(context.Background(), &ret, upToken, key, src, file.Size, &putExtra)

	// 以默认key方式上传
	// err = formUploader.PutWithoutKey(context.Background(), &ret, upToken, src, fileSize, &putExtra)

	// 自定义key，上传指定路径的文件
	// localFilePath = "./aa.jpg"
	// err = formUploader.PutFile(context.Background(), &ret, upToken, key, localFilePath, &putExtra)

	// 默认key，上传指定路径的文件
	localFilePath := "./aa.jpg"
	err = formUploader.PutFile(context.Background(), &ret, upToken, key, localFilePath, &putExtra)

	if err != nil {
		code := 501
		return code, err.Error()
	}

	url := imgUrl + ret.PersistentID// 返回上传后的文件访问路径
	return 0, url



}
