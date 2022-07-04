package model

import "github.com/jinzhu/gorm"

type Admin struct {
	*Model
	Name string
	State uint8
}

func (a Admin) TableName() string {
	return "cms_admin"
}

func (a Admin) Count(db *gorm.DB) (int, error) {
	var count int
	if a.Name != "" {
		db = db.Where("name = ?", a.Name)
	}
	db = db.Where("state = ?", a.State)
	if err := db.Model(&a).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (a Admin) List(db *gorm.DB, pageOffset, pageSize int) ([]*Admin, error) {
	var admins []*Admin
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if a.Name != "" {
		db = db.Where("name = ?", a.Name)
	}

	db = db.Where("state = ?", a.State)
	if err = db.Where("is_del = ?", 0).Find(&admins).Error; err != nil {
		return nil, err
	}
	return admins, nil
}

func (a Admin) Create (db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Admin) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&a).Where("id = ? AND is_del = ?", a.ID, 0).Update(values).Error
}

func (a Admin) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ? ", a.Model.ID, 0).Delete(&a).Error
}

