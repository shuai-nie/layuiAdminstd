package service

import (
	"layuiAdminstd/pkg/app"
)

type AuthGroupListRequest struct {
	Module string
	Type uint8
	Title string
	Description string
	Status uint8
}

type CreateAuthGroupRequest struct {
	Module      string `form:"module"`
	Type        uint8  `form:"type"`
	Title       string `form:"title"`
	Description string `form:"description"`
	Status      uint32 `form:"status"`
	Rules       string `form:"rules"`
}

type UpdateAuthGroupRequest struct {
	ID          uint32 `form:"id" binding:"required,gte=1"`
	Module      string `form:"module"`
	Type        uint8  `form:"type"`
	Title       string `form:"title"`
	Description string `form:"description"`
	Status      uint32 `form:"status"`
	Rules       string `form:"rules"`
}

type AuthGroup struct {
	Id          uint32 `form:"id"`
	Module      string `form:"module"`
	Type        uint8  `form:"type"`
	Title       string `form:"title"`
	Description string `form:"description"`
	Status      uint8  `form:"status"`
}

func (svc *Service) GetAuthGroupList(param *AuthGroupListRequest, pager *app.Pager) ([]*AuthGroup, int, error) {
	adminCount, err := svc.dao.CountAdminList(param.Status)
	var adminList []*AuthGroup
	admins, err := svc.dao.GetAuthGroupList(param.Status, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}
	for _, admin := range admins {
		adminList = append(adminList, &AuthGroup{
			Id: admin.ID,
			Title: admin.Title,
			Module: admin.Module,
			Description: admin.Description,
			Type: admin.Type,
		})
	}

	return adminList, adminCount, nil
}

func (svc *Service) CreateAuthGroup (param *CreateAuthGroupRequest) error {
	_, err := svc.dao.CreateAuthGroup(1,"2","3","4")

	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) UpdateAuthGroup (param *UpdateAuthGroupRequest) error {
	return svc.dao.UpdateAuthGroup(1, 1, "", "", "")
}