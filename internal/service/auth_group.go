package service

import (
	"layuiAdminstd/internal/dao"
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

type AuthGroupRequest struct {
	Id          uint32 `form:"id"`
	Module      string `form:"module"`
	Type        uint8  `form:"type"`
	Title       string `form:"title"`
	Description string `form:"description"`
	Status      uint8  `form:"status"`
}

func (svc *Service) GetAuthGroup(param *AuthGroupRequest) (*AuthGroup, error) {
	authgroup, err := svc.dao.GetAuthGroup(param.Id, param.Status)
	if err != nil {
		return nil, err
	}
	return &AuthGroup{
		Id: authgroup.ID,
		Module: authgroup.Module,
		Title: authgroup.Title,
		Description: authgroup.Description,
		Status: authgroup.Status,
	}, nil
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
	_, err := svc.dao.CreateAuthGroup(&dao.AuthGroup{
		Module: param.Module,
		Type: param.Type,
		Title: param.Title,
		Description: param.Description,
		Status: param.Status,
		Rules: param.Rules,
	})

	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) UpdateAuthGroup (param *UpdateAuthGroupRequest) error {
	err := svc.dao.UpdateAuthGroup(&dao.AuthGroup{
		ID: param.ID,
		Title: param.Title,
	})

	if err != nil {
		return err
	}
	return nil
}