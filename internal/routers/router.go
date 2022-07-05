package routers

import (
	"github.com/gin-gonic/gin"
	"layuiAdminstd/internal/routers/api"
	v1 "layuiAdminstd/internal/routers/api/v1"
)


func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	admin := v1.NewAdmin()
	upload := api.NewUpload()

	// 上传文件 url
	r.POST("/upload/file", upload.UploadFile)

	apiv1 := r.Group("/api/v1")
	{
		// 创建
		apiv1.POST("/admin", admin.Create)
		// 删除
		apiv1.DELETE("/admins/:id", admin.Delete)
		// 更新
		apiv1.PUT("/admins/", admin.Update)
		// 获取列表
		apiv1.GET("/admins", admin.List)

		apiv1.POST("/tags")
		apiv1.DELETE("/tags/:id")
		apiv1.PUT("/tags/:id")
		apiv1.PATCH("/tags/:id/state")
		apiv1.GET("/tags")

		apiv1.POST("/articles")
		apiv1.DELETE("/articles/:id")
		apiv1.PUT("/articles/:id")
		apiv1.PATCH("articles/:id/state")
		apiv1.GET("/articles/:id")
		apiv1.GET("/articles")
	}
	return r
}
