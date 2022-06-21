package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UsersController struct {

}

func (controller *UsersController) Get(context *gin.Context) {
	id := context.Query("id")
	context.JSON(http.StatusOK, gin.H{
		"id" : id,
		"config" : config.GetConfig(),
	})
}
