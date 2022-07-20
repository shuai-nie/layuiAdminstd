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
	Module string
	Type uint8
	Title string
	Description string
	Status uint32
	Rules string
}

type AuthGroup struct {
	Id uint32
	Module string
	Type uint8
	Title string
	Description string
	Status uint8
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
	})

	if err != nil {
		return err
	}
	return nil
}