package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"layuiAdminstd/global"
	"layuiAdminstd/internal/routers/api"
	"layuiAdminstd/internal/routers/api/v1"
	"net/http"
	"strings"
)


func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger()) // 日志
	r.Use(Cors()) // 跨域请求

	r.Use(gin.Recovery())
	gin.SetMode(global.ServerSetting.RunMode)

	admin := v1.NewAdmin()
	AuthGroup := v1.NewAuthGroup()
	menu := v1.NewMenu()
	upload := api.NewUpload()

	// 上传文件 url
	r.POST("/api/upload/file", upload.UploadFile)

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
		apiv1.GET("/admins/:id", admin.Get)

		// 权限角色列表
		apiv1.GET("/auth/group", AuthGroup.List)
		apiv1.POST("/auth/group/create", AuthGroup.Create)
		apiv1.GET("/auth/group/:id", AuthGroup.Get)
		apiv1.PUT("/auth/group", AuthGroup.Update)

		apiv1.GET("/menu/index", menu.List)


/*		apiv1.POST("/tags")
		apiv1.DELETE("/tags/:id")
		apiv1.PUT("/tags/:id")
		apiv1.PATCH("/tags/:id/state")
		apiv1.GET("/tags")

		apiv1.POST("/articles")
		apiv1.DELETE("/articles/:id")
		apiv1.PUT("/articles/:id")
		apiv1.PATCH("articles/:id/state")
		apiv1.GET("/articles/:id")
		apiv1.GET("/articles")*/
	}
	return r
}

// 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method	// 请求方法
		origin := c.Request.Header.Get("Origin")	// 请求头部
		var headerKeys []string		// 声明请求头 keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-all-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-ogigin, access-control-allow-headers"
		}

		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")	// 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE, UPDATE") // 服务器支持的所有跨域请求的方法，为了避免浏览次数请求的多次`预检`请求
			// header 的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			// 允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")
		}

		// 放行所有options 方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()

	}
}