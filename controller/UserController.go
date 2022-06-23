package controller

import (
	"github.com/gin-gonic/gin"
	"layuiAdminstd/model"
	"net/http"
	"strconv"
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

// 查找用户
func FindUserDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(200, gin.H{
			"data": "请上传正确的数据",
		})
		return
	}
	result := model.FindForId(id)
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}

// 编辑用户
func UpdateUserInfo(c *gin.Context) {
	//name := c.Param("name")
	//jwt := c.Request.Header.Get("jwt")
	//name := c.Query("name")
	// 查询
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data":"请选择正确的数据",
		})
		return
	}
	result := model.FindForId(id)
	if result.Id != id {
		c.JSON(http.StatusOK, gin.H{
			"data" : "用户不存在",
		})
		return
	}
	//sex, _ := strconv.Atoi(c.PostForm("sex"))
	newInfo := map[string]interface{}{
		"username" : c.PostForm("username"),
	}
	updateRes := model.UpdateUserForId(id, newInfo)
	c.JSON(http.StatusOK, gin.H{
		"data" : updateRes,
	})
}

func DeleteUserInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data":"请选择正确的数据",
		})
		return
	}
	result := model.FindForId(id)
	if result.Id != id {
		c.JSON(http.StatusOK, gin.H{
			"data":"用户不存在",
		})
		return
	}
	res := model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

//https://blog.csdn.net/qq_42067682/article/details/123089750