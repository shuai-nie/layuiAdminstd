package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"layuiAdminstd/global"
	"layuiAdminstd/internal/service"
	"layuiAdminstd/pkg/app"
	"layuiAdminstd/pkg/errcode"
)

type AuthGroup struct {}

func NewAuthGroup() AuthGroup {
	return AuthGroup{}
}

func (a AuthGroup) Get(c *gin.Context) {}

func (a AuthGroup) List(c *gin.Context) {
	param := service.AuthGroupListRequest{}
	response := app.NewResponse(c)

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}

	fmt.Println(pager)

	admins, totalRows, err := svc.GetAuthGroupList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetAdminList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetAdminListFail)
		return
	}
	response.ToResponseList(admins, totalRows)
	return
}

func (a AuthGroup) Create (c *gin.Context) {
	param := service.CreateAuthGroupRequest{}
	response := app.NewResponse(c)
	//  c.Request
	svc := service.New(c.Request.Context())
	err := svc.CreateAuthGroup(&param)
	if err != nil {
		return
	}

	response.ToResponse(gin.H{
		"code" : 0,
		"message": "成功",
	})
	return
}
