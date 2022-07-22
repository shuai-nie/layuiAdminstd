package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"layuiAdminstd/global"
	"layuiAdminstd/internal/service"
	"layuiAdminstd/pkg/app"
	"layuiAdminstd/pkg/convert"
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
	vaild, errs := app.BindAndValid(c, &param)
	if !vaild {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CreateAuthGroup(&param)
	if err != nil {
		global.Logger.Errorf("svc.createAuthGroup err %v", err)
		response.ToErrorResponse(errcode.ErrorCreateAuthGroupFail)
		return
	}

	response.ToResponse(gin.H{
		"code" : 0,
		"message": "成功",
		"data": param.Module,
	})
	return
}

func (a AuthGroup) Update (c *gin.Context) {
	param := service.UpdateAuthGroupRequest{ID:convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateAuthGroup(&param)
	if err != nil {
		global.Logger.Errorf(" svc.UpdateAuthor err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateAuthGroupRequestFaill)
		return
	}

	response.ToResponse(gin.H{
		"code" : 0,
		"message": "成功",
	})
	return
}
