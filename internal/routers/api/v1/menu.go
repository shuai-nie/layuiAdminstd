package v1

import (
	"github.com/gin-gonic/gin"
	"layuiAdminstd/global"
	"layuiAdminstd/internal/service"
	"layuiAdminstd/pkg/app"
	"layuiAdminstd/pkg/convert"
	"layuiAdminstd/pkg/errcode"
)

type Menu struct {}

func NewMenu() Menu {
	return Menu{}
}

func (a Menu) Get(c *gin.Context) {
	param := service.MenuRequest{Id: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	auth, err := svc.GetAuthGroup(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetAuthGroup err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetAuthGroupRequestFail)
		return
	}
	response.ToResponse(auth)
	return
}