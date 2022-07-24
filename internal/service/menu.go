package service

type MenuRequest struct {
	Id uint32
}

func (svc *Service) GetMenu(param *MenuRequest) (*MenuRequest, error) {
	menu, err := svc.dao.GetMenu(param.Id)
	if err != nil {
		return nil, err
	}
	return &MenuRequest{
		Id: menu.ID,
	}, nil
}
