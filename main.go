package main

import (
	"bytes"
	"fmt"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"io/ioutil"
	"layuiAdminstd/config"
	"layuiAdminstd/model"
	"layuiAdminstd/pkg/logger"
	"net/http"
	"time"
)

var (
	Logger *logger.Logger
)

func main() {
	conf, err := config.ParseConfig("./config/app.json")
	if err != nil {
		panic("读取配置文件失败，" + err.Error())
	}

	fmt.Println("config:%v\n", conf.AppHost)

	//fmt.Println("config:#v\n" , conf)

	// 创建一个不包含任何中间件的engine
	r := gin.New()
	// 添加日志中间件
	r.Use(gin.Logger())
	// 添加错误回复重定向中间件
	r.Use(gin.Recovery())
	// 以上三部可以使用r := gin.Default() 一步实现
	// 对 /benchmark路由添加两个处理函数
	//r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	r.GET("/login/:username/:password", login)
	r.GET("/verify/:token", verify)
	r.GET("/refresh/:token", refresh)
	r.GET("/sayHello/:token", sayHello)

	r.GET("/", )

	// 链接数据库
	model.OpenDB()
	// 设置连接池
	model.SetPool()
	// 关闭数据库
	defer model.CloseDB()



	// 创建路由分组
	authorizd := r.Group("/v1")
	// 使用 authRequired 中间件
	authorizd.Use(AuthRequired())
	{
		/*authorizd.POST("/login", loginEndopint)
		authorizd.POST("/submit", submitEndpoint)
		authorizd.POST("/read", readEndpoint)

		testing := authorizd.Group("testing")
		testing.GET("/analytics", anlyticsEndpoint)*/

	}

	r.Run()
}

func setupLogger() error {
	Logger = logger.NewLogger(&lumberjack.Lo)

}

// 中间件
func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := ctx.GetRawData()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("data: %v\n", string(data))

		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		ctx.Next()
	}
}

func AuthRequired() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"foo":"bar",
		"austin":"1234",
		"lena" : "hell2",
		"menu" : "4321",
	})
}

const (
	ErrorReason_ServerBusy = "服务器繁忙"
	ErrorReason_ReLogin = "请重新登录"
)

func sayHello(c *gin.Context) {
	strToken := c.Param("token")
	claim, err := verifyAction(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, "hello", claim.Username)
}

type JWTClaims struct {
	jwt.StandardClaims
	UserID      int      `json:"user_id"`
	Password    string   `json:"password"`
	Username    string   `json:"username"`
	FullName    string   `json:"full_name"`
	Permissions []string `json:"permissions"`
}

var (
	Secret     = "dong_tech" // 加盐
	ExpireTime = 3600        // token有效期
)

func login(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")
	claims := &JWTClaims{
		UserID:      1,
		Username:    username,
		Password:    password,
		FullName:    username,
		Permissions: []string{},
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	signedToken,err:=getToken(claims)
	if err!=nil{
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, signedToken)
}

func verify(c *gin.Context) {
	strToken := c.Param("token")
	claim,err := verifyAction(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, "verify,", claim.Username)
}

func refresh(c *gin.Context) {
	strToken := c.Param("token")
	claims,err := verifyAction(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
	signedToken,err:=getToken(claims)
	if err!=nil{
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, signedToken)
}

func verifyAction(strToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, errors.New(ErrorReason_ServerBusy)
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	fmt.Println("verify")
	return claims, nil
}

func getToken(claims *JWTClaims)(string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "",errors.New(ErrorReason_ServerBusy)
	}
	return signedToken,nil
}