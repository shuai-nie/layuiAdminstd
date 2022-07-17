package app

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"layuiAdminstd/global"
	"layuiAdminstd/pkg/util"
	"time"
)

type Claims struct {
	AppKey string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

func GetJWTSecret() string  {
	return global.JWTSetting.Secret
}

// 主要功能是生成JWT token
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey: util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: global.JWTSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

// 主要功能是解析和校验 token
func ParseToken(token string) (*Claims, error) {
	// ParseWithClaims 用于解析鉴权的声明
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error ) {
		return GetJWTSecret(), nil
	})

	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		// valid 验证基于时间的声明
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}


func test(c *gin.Context) {
	// 加密
	mySigningKey := []byte("    ")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":"",
		"id":"",
		"exp":time.Now().Unix()+5,
		"iss":"sywdebug",
	})
	fmt.Println(token)

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return
	}

	fmt.Println("加密后的token 字符串", tokenString)

	// 解密
	token, err = jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return
	}
	fmt.Println("token:", token)
	fmt.Println("token.Claims:", token.Claims)
	fmt.Println(token.Claims.(jwt.MapClaims)["name"])
}