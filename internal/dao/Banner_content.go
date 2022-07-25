package dao

import (
"layuiAdminstd/internal/model"
"layuiAdminstd/pkg/app"
)

type BannerContent struct {
	ID uint32
	Module string
	Type uint8
	Title string
	Description string
	Status uint32
	Rules string
}

func (d *Dao) CreateBannerContent(param *BannerContent) (*model.BannerContent, error) {
	admin := model.BannerContent{
		Module: param.Module,
		Type: param.Type,
		Title: param.Title,
		Description: param.Description,
	}
	return admin.Create(d.engine)
}

func (d *Dao) UpdateBannerContent(param *BannerContent) error {
	admin := model.BannerContent{Model:&model.Model{ID: param.ID}}
	values := map[string]interface{}{
		"title": param.Title,
		"description": param.Description,
	}
	return admin.Update(d.engine, values)
}

func (d *Dao) GetBannerContent(id uint32, status uint8) (model.BannerContent, error) {
	admin := model.BannerContent{Model: &model.Model{ID: id}, Status:status}
	return admin.Get(d.engine)
}

//统计条数
func (d *Dao) CountBannerContentList(status uint8) (int, error) {
	admin := model.AuthGroup{Status: status}
	return admin.Count(d.engine)
}

func (d *Dao) GetBannerContentList(status uint8, page, pageSize int) ([]*model.BannerContent, error) {
	admin := model.BannerContent{Status: status}
	return admin.List(d.engine, app.GetPageOffset(page, pageSize), pageSize)
}
