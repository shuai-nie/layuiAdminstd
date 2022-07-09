package service

import (
	"layuiAdminstd/internal/dao"
	"layuiAdminstd/pkg/app"
)

type CountAdminRequest struct {
	State uint8
}

type AdminRequest struct {
	ID uint32
	State uint8
}

type AdminListRequest struct {
	Name string
	Status uint8
}

type Admin struct {
	ID uint32
	Name string
	Username string
	Email string
	GroupId uint32
	str string
	Status uint8
	// 未完成
}

type CreateAdminRequest struct {
	Name string
	CreateBy string
	State uint8
}

type UpdateAdminRequest struct {
	ID uint32
	Name string
	State uint8
	ModifiedBy string
}

type DeleteAdminRequest struct {
	ID uint32
}

func (svc *Service) GetAdmin (param *AdminRequest) (*Admin, error) {
	admin, err := svc.dao.GetAdmin(param.ID, param.State)
	if err != nil {
		return nil, err
	}
	return &Admin{
		ID: admin.ID,
		// 未完成
	}, nil
}

func (svc *Service) GetAdminList(param *AdminListRequest, pager *app.Pager) ([]*Admin, int, error) {
	adminCount, err := svc.dao.CountAdminList(param.Status)
	var adminList []*Admin
	admins, err := svc.dao.GetAdminList(param.Status, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}
	for _, admin := range admins {
		adminList = append(adminList, &Admin{
			ID: admin.ID,
			Name: admin.Name,
			Username: admin.Username,
			Email: admin.Email,
			GroupId: admin.GroupId,
		})
	}

	return adminList, adminCount, nil
}

func (svc *Service) CreateAdmin (param *CreateAdminRequest) error {
	_, err := svc.dao.CreateAdmin(&dao.Admin{
		Name: param.Name,
	})

	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) UpdateAdmin (param *UpdateAdminRequest) error {
	err := svc.dao.UpdateAdmin(&dao.Admin{
		ID: param.ID,
		Name: param.Name,
	})

	if err != nil {
		return err
	}

	return nil
}

