package model

import (
"github.com/jinzhu/gorm"
)


type BannerContent struct {
	*Model
	Module string
	Type uint8
	Title string
	Description string
	Status uint8
	Rules string
}

func (a BannerContent) TableName() string {
	return "jm_banner_content"
}

func (a BannerContent) Create(db *gorm.DB) (*BannerContent, error) {
	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (a BannerContent) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&a).Where("id = ?", a.ID).Update(values).Error
}

func (a BannerContent) Get(db *gorm.DB) (BannerContent, error) {
	var auth_group BannerContent
	if err := db.Where("id = ?", a.ID).First(&auth_group).Error; err != nil {
		return auth_group, err
	}
	return auth_group, nil
}

func (a BannerContent) Count(db *gorm.DB) (int, error) {
	var count int
	if err := db.Model(&a).Where("status = 1").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (a BannerContent) List(db *gorm.DB, pageOffset, pageSize int) ([]*BannerContent, error) {
	var BannerContent []*BannerContent
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if err = db.Where("status = 1").Find(&BannerContent).Error; err != nil {
		return nil, err
	}
	return BannerContent, nil

}

func (a BannerContent) Delete(db *gorm.DB) error {
	return db.Where("id = ? ", a.ID).Update("status=0").Error
}


/*func (a AuthGroup) Create(db *gorm.DB) (*AuthGroup, error) {}
func (a AuthGroup) Update(db *gorm.DB) error {}
func (a AuthGroup) Get(db *gorm.DB) (AuthGroup, error) {}
func (a AuthGroup) Count(db *gorm.DB) (int, error) {}
func (a AuthGroup) List(db *gorm.DB) ([]*AuthGroup, error) {}
func (a AuthGroup) Delete(db *gorm.DB) error {}*/