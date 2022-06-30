package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//var (
//	Logger *logger.Logger
//)

func InitDB() *gorm.DB {
	dbConfig := GetConfig().Database
	driverName := "mysql"
	host := dbConfig.Host
	port := dbConfig.Port
	database := dbConfig.DbName
	username := dbConfig.User
	password := dbConfig.Password
	charset := dbConfig.Charset
	//fmt.Println("config:  ", host, port, database, username, password, charset)
	//Logger.Infof("%s ===  %s", host, port)
	//panic(host + port + database + username + password + charset)

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