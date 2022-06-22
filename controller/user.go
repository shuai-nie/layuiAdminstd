package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
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

/************/
type UserController struct {
}
var UserControl = &UserController{}

func (t *UserController) RoleList(c *gin.Context) {
	// 解析前端参数
	params, pageNo, pageSize := PackageRequestParams(c)
	res, total := Service.UserServer.RoleList(params, pageNo, pageSize)
	ResponseSuccessMap(c, map[string]interface{}{
		"records": res,
		"total":   total,
	})
}

func (t *UserController) RoleEdit(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	CheckError(err)
	var msg Model.UserRole
	json.Unmarshal(data, &msg)
	user, _ := GetCurUserinfo(c)
	// time.Time 与数据库与前端三者之间转换实在太麻烦，所以定义为string, updateTime为空时指定一个特定时间，前端显示时过滤掉
	if msg.Id == 0 {
		msg.CreateBy = user.Username
		msg.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		msg.UpdateTime = "1010-10-10 00:00:00"
	} else {
		msg.UpdateBy = user.Username
		msg.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	}
	if err := Service.UserServer.RoleEdit(msg, msg.Id != 0); err == nil {
		ResponseSuccess(c, "success")
	} else {
		ResponseError(c, err.Error())
	}
}

func (t *UserController) RoleDel(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	CheckError(err)
	var msg struct {
		Ids []int
	}
	json.Unmarshal(data, &msg)
	if err := Service.UserServer.RoleDel(msg.Ids); err == nil {
		ResponseSuccess(c, "success")
	} else {
		ResponseError(c, err.Error())
	}
}