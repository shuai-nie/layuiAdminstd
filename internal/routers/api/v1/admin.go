package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"layuiAdminstd/internal/service"
	"layuiAdminstd/pkg/app"
)

type Admin struct {}

func NewAdmin() Admin {
	return Admin{}
}

func (a Admin) Get(c *gin.Context) {}

// 查询列表
func (a Admin) List(c *gin.Context) {
	fmt.Print("ssssss")

	param := service.AdminListRequest{}
	response := app.NewResponse(c)

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	admins, totalRows, err := svc.GetAdminList(&param, &pager)
	if err != nil {
		return
	}
	response.ToResponseList(admins, totalRows)
	return
}
func (a Admin) Create(c *gin.Context) {}
func (a Admin) Update(c *gin.Context) {}
func (a Admin) Delete(c *gin.Context) {}
