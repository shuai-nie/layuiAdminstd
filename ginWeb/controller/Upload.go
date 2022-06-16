package controller

import (
	"github.com/gin-gonic/gin"
	"ginWeb/lib/func"

)


type UploadController struct {}

var UploadControl = &UploadController{}

func (*UploadController) UploadImage(c *gin.Context) {
	dir := Config.StaticUrl // 配置文件中的静态资源路径
	file, _ := c.FormFile("file")
	filename := RandStr(40) + ".jpg"
	//将文件保存到本地服务器的指定位置
	if err := c.SaveUploadedFile(file, dir+filename); err != nil {
		ResponseError(c, "文件上传失败")
	} else {
		ResponseSuccess(c, filename)
	}
}
// 某些文件的执行依赖于其他文件，比如3D模型，所以需要通过ZIP的方式上传，同时要指定执行文件的名称
// 先将ZIP保存，然后解压，最后删除保存的ZIP
func (*UploadController) UploadZipFile(c *gin.Context) {
	dir := Config.StaticUrl
	file, _ := c.FormFile("file")
	name := c.PostForm("name") // 执行文件名称
	filename := RandStr(40)
	//将文件保存到本地服务器的指定位置
	if err := c.SaveUploadedFile(file, dir+filename+".zip"); err != nil {
		ResponseError(c, "文件上传失败")
	} else {
		det, err := Unzip(filename+".zip", filename, name) // 解压
		if err != nil {
			ResponseError(c, "文件上传失败")
			return
		}
		UploadControl.RemoveFile(filename + ".zip") // 删除zip
		if det == "" {
			ResponseError(c, "文件名称错误")
		} else {
			ResponseSuccess(c, det)
		}
	}
}