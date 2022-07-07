package dao

import (
	"layuiAdminstd/internal/model"
	"layuiAdminstd/pkg/app"
)

type Admin struct {
	ID uint32
	Name string
	Password string
	email string
	GroupId uint32
	Str string
	Status uint8
}

func (d *Dao) CreateAdmin(param *Admin) (*model.Admin, error) {
	admin := model.Admin{
	}
	return admin.Create(d.engine)
}


func (d *Dao) UpdateAdmin(param *Admin) error {
	admin := model.Admin{Model:&model.Model{ID: param.ID}}
	values := map[string]interface{}{
		// 数据格式
	}
	return admin.Update(d.engine, values)
}

func (d *Dao) GetAdmin(id uint32, state uint8) (model.Admin, error) {
	admin := model.Admin{Model: &model.Model{ID: id}, State:state}
	return admin.Get(d.engine)
}

//统计条数
func (d *Dao) CountAdminList(state uint8) (int, error) {
	admin := model.Admin{State: state}
	return admin.Count(d.engine)
}

func (d *Dao) GetAdminList(state uint8, page, pageSize int) ([]*model.Admin, error) {
	admin := model.Admin{State: state}
	return admin.List(d.engine, app.GetPageOffset(page, pageSize), pageSize)
}