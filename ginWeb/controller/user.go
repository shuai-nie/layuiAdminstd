package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type UserController struct {}

func (t *UserController) Login(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	CheckError(err)
	var msg struct {
		UserName string
		Password string
	}
	json.Unmarshal(data, &msg)
	res, _ := Service.UserServer.UserList(fmt.Sprintf("username = '%s'", msg.UserName), 1, 1)
	if len(res) > 0 && res[0].Password == msg.Password {
		// 生成随机字符串
		str := RandStr(32)
		result, _ := json.Marshal(res[0])
		// 将用户信息存入redis中，设置过期时间为10分钟
		Redis.SetRedis(str, string(result), 60*10)
		ResponseSuccessMap(c, map[string]interface{}{
			"token":    str,
			"userinfo": res[0],
		})
	} else {
		ResponseError(c, "error5")
	}
}