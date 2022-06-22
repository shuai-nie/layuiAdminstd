package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

type RequestCondition struct {
	Column string
	Value  string
	Type   string
}
// 前端列表请求参数格式
type RequestListData struct {
	PageNo    int
	PageSize  int
	Condition []RequestCondition
}
func ResponseSuccess(c *gin.Context, data string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    data,
		"success": true,
	})
}
func ResponseSuccessMap(c *gin.Context, data map[string]interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    data,
		"success": true,
	})
}
func ResponseError(c *gin.Context, data string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    "500",
		"message": data,
		"data":    "",
		"success": false,
	})
}
func CheckError(e error) {
	if e != nil {
		fmt.Println(e.Error())
	}
}
// 将前端请求参数转换为 grom where语句
func PackageRequestParams(c *gin.Context) (string, int, int) {
	data, err := ioutil.ReadAll(c.Request.Body)
	CheckError(err)
	var req RequestListData
	json.Unmarshal(data, &req)
	var res []string
	for _, value := range req.Condition {
		var c string
		key := snakeString(value.Column)
		switch value.Type {
		case "eq":
			c = key + " = " + "'" + value.Value + "'"
			break
		case "like":
			c = key + " like " + "'%" + value.Value + "%'"
			break
		case "gt":
			c = key + " > " + value.Value
			break
		case "lt":
			c = key + " < " + value.Value
			break
		}
		res = append(res, c)
	}
	return strings.Join(res, " and "), req.PageNo, req.PageSize
}


// 驼峰转蛇形
func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}