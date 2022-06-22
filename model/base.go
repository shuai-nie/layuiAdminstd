package model

import (
	"github.com/jinzhu/gorm"
	"layuiAdminstd/config"
	"time"
)

var db *gorm.DB

func SetPool() {
	sqlDB := GetDB().DB()
	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxIdleConns(100)

	sqlDB.SetConnMaxLifetime(time.Hour)

}

// OpenDB 开启数据库
func OpenDB() {
	db = config.InitDB()
}

// GetDB 链接数据库
func GetDB() *gorm.DB {
	return db
}

// CloseDB 关闭数据库
func CloseDB() {
	db.Close();
}