package dao

import "layuiAdminstd/internal/model"

type Menu struct {
	ID uint32
	Module string
	Type uint8
	Title string
	Description string
	Status uint32
	Rules string
}

func (d *Dao) GetMenu(id uint32) (model.Menu, error) {
	admin := model.Menu{Model: &model.Model{ID: id}}
	return admin.Get(d.engine)
}