package api

import (
	"github.com/gin-gonic/gin"
	"layuiAdminstd/global"
	"layuiAdminstd/internal/service"
	"layuiAdminstd/pkg/app"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app %v", errs)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})


}