package dao

import (
	"fmt"
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

func (d *Dao) CreateAuthGroup(param *AuthGroup) (*model.AuthGroup, error) {

	admin := model.AuthGroup{
		Module: param.Module,
		Type: param.Type,
		Title: param.Title,
		Description: param.Description,
	}
	fmt.Println("+++++++++++++++++++++++")
	fmt.Println(param.Module)
	fmt.Println("+++++++++++++++++++++++")

	return admin.Create(d.engine)
}

func (d *Dao) UpdateAuthGroup(param *AuthGroup) error {
	admin := model.AuthGroup{Model:&model.Model{ID: param.ID}}
	values := map[string]interface{}{
		// 数据格式
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
