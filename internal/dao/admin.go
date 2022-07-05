package dao

import "layuiAdminstd/internal/model"

type Admin struct {
	ID uint32
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