package v1

import (
	"blog-service/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"layuiAdminstd/internal/service"
	"layuiAdminstd/pkg/app"
	"layuiAdminstd/pkg/convert"
	"layuiAdminstd/pkg/errcode"
)

type Admin struct {}

func NewAdmin() Admin {
	return Admin{}
}

func (a Admin) Get(c *gin.Context) {}

// @Summary 获取多个
// @Produce json
// @Param name query string false "文章名称"
// @Param tag_id query int false "标签ID"
// @Param state query int false "状态"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [get]
func (a Admin) List(c *gin.Context) {
	fmt.Print("ssssss")

	param := service.AdminListRequest{}
	response := app.NewResponse(c)

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}

	fmt.Println(pager)

	admins, totalRows, err := svc.GetAdminList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetAdminList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetAdminListFail)
		return
	}
	response.ToResponseList(admins, totalRows)
	return
}

// @Summary 创建
// @Produce json
// @Param tag_id body string true "标签ID"
// @Param title body string true "文章标题"
// @Param desc body string false "文章简述"
// @Param cover_image_url body string true "封面图片地址"
// @Param content body string true "文章内容"
// @Param created_by body int true "创建者"
// @Param state body int false "状态"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post]
func (a Admin) Create(c *gin.Context) {
	param := service.CreateAdminRequest{}
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	err := svc.CreateAdmin(&param)
	if err != nil {
		global.Logger.Errorf( "svc.CreateAdmin err : %v", err)
		response.ToErrorResponse(errcode.ErrorCreateAdminFail)
		return
	}
}

func (a Admin) Update(c *gin.Context) {
	param := service.UpdateAdminRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	err := svc.UpdateAdmin(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateAdmin err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateAdminFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

func (a Admin) Delete(c *gin.Context) {}

