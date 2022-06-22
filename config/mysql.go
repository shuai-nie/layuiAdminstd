package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	dbConfig := GetConfig().Database
	driverName := "mysql"
	host := dbConfig.Host
	port := dbConfig.Port
	database := dbConfig.DbName
	username := dbConfig.User
	password := dbConfig.Password
	charset := dbConfig.Chartset

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}

	return db
}