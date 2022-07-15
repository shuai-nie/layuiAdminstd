package api

import (
	"github.com/gin-gonic/gin"
	"layuiAdminstd/internal/service"
	"layuiAdminstd/pkg/app"
	"layuiAdminstd/pkg/convert"
	"layuiAdminstd/pkg/upload"
	"net/http"
	"os"
	"path"
)

type Upload struct {}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		return 
	}
	
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		return
	}
	
	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}

// 下载次数
func fileDownload(c *gin.Context) {
	filePath := c.Query("url")
	// 打开文件
	fileTmp, errByOpenFile := os.Open(filePath)
	defer fileTmp.Close()

	// 获取文件的名称
	fileName := path.Base(filePath)
	if errByOpenFile != nil {
		c.JSON(http.StatusOK, gin.H{
			"message":"文件获取失败",
		})
	}

	c.Header("Content-Type", "application/octet-stream")
	// 强制浏览器下载
	c.Header("Content-Disposition", "attachment; filename=" + fileName)
	// 浏览器下载或预览
	c.Header("Content-Disposition", "inline; filename=" + fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	c.File(filePath)
	return
}

