package model

type Admin struct {
	*Model
	Name string
	State uint8
}

func (a Admin) TableName() string {
	return "cms_admin"
}

