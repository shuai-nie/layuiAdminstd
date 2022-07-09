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

func (d *Dao) GetAdmin(id uint32, status uint8) (model.Admin, error) {
	admin := model.Admin{Model: &model.Model{ID: id}, Status:status}
	return admin.Get(d.engine)
}

//统计条数
func (d *Dao) CountAdminList(status uint8) (int, error) {
	admin := model.Admin{Status: status}
	return admin.Count(d.engine)
}

func (d *Dao) GetAdminList(status uint8, page, pageSize int) ([]*model.Admin, error) {
	admin := model.Admin{Status: status}
	return admin.List(d.engine, app.GetPageOffset(page, pageSize), pageSize)
}