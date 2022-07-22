package dao

import (
	"layuiAdminstd/internal/model"
	"layuiAdminstd/pkg/app"
)

type AuthGroup struct {
	ID uint32
	Module string
	Type uint8
	Title string
	Description string
	Status uint32
	Rules string
}

func (d *Dao) CreateAuthGroup(types int8, module, title, desciption string) (*model.AuthGroup, error) {

	admin := model.AuthGroup{
		Module: module,
		Title: title,
		Description: desciption,
	}
	return admin.Create(d.engine)
}

func (d *Dao) UpdateAuthGroup(id uint32, types int8, module, title, description string) error {
	admin := model.AuthGroup{Model:&model.Model{ID: id}}
	values := map[string]interface{}{
		"Module": module,
		"Type": types,
		"Title": title,
		"Description": description,
	}
	return admin.Update(d.engine, values)
}

func (d *Dao) GetAuthGroup(id uint32, status uint8) (model.AuthGroup, error) {
	admin := model.AuthGroup{Model: &model.Model{ID: id}, Status:status}
	return admin.Get(d.engine)
}

//统计条数
func (d *Dao) CountAuthGroupList(status uint8) (int, error) {
	admin := model.AuthGroup{Status: status}
	return admin.Count(d.engine)
}

func (d *Dao) GetAuthGroupList(status uint8, page, pageSize int) ([]*model.AuthGroup, error) {
	admin := model.AuthGroup{Status: status}
	return admin.List(d.engine, app.GetPageOffset(page, pageSize), pageSize)
}
