package model

import "github.com/jinzhu/gorm"

type Menu struct {
	*Model
}

func (a Menu) TableName() string {
	return "jm_auth_rule"
}

func (a Menu) Get(db *gorm.DB) (Menu, error) {
	var menu Menu
	if err := db.Where("id = ?", a.ID).First(&menu).Error; err != nil {
		return menu, err
	}
	return menu, nil
}