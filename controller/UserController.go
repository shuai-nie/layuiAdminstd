package controller

import (
	"github.com/gin-gonic/gin"
	"layuiAdminstd/model"
	"net/http"
)

func TestMysql(c *gin.Context) {
	db := model.GetDB()
	var user model.User
	db.Table("tpsw_user").First(&user)
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func CreateNewUser(c *gin.Context) {
	//sex, _ := strconv.Atoi(c.PostForm("sex"))
	user := model.User{
		// 获取post 传输数据
		Username: c.PostForm("username"),
	}

	result := model.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}

//https://blog.csdn.net/qq_42067682/article/details/123089750