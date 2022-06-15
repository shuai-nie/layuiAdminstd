package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建一个不包含任何中间件的engine
	r := gin.New()
	// 添加日志中间件
	r.Use(gin.Logger())
	// 添加错误回复重定向中间件
	r.Use(gin.Recovery())
	// 以上三部可以使用r := gin.Default() 一步实现
	// 对 /benchmark路由添加两个处理函数
	r.GET("/benchmark", MyBenchLogger(), benchEndpoint)
	// 创建路由分组
	authorizd := r.Group("/")
	// 使用 authRequired 中间件
	authorizd.Use(AuthRequired())
	{

		authorizd.POST("/login", loginEndopint)
		authorizd.POST("/submit", submitEndpoint)
		authorizd.POST("/read", readEndpoint)

		testing := authorizd.Group("testing")
		testing.GET("/analytics", anlyticsEndpoint)

	}

	r.Run()
}

func AuthRequired() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"foo":"bar",
		"austin":"1234",
		"lena" : "hell2",
		"menu" : "4321",
	})
}