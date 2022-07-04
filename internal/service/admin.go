package service

type CountAdminRequest struct {
	State uint8
}

type AdminListRequest struct {
	Name string
	State uint8
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
