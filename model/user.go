package model

import "github.com/jinzhu/gorm"

type User struct {
	Id         int    `gorm:"column:id" db:"id" json:"id" form:"id"`
	Username       string `gorm:"column:username" db:"username" json:"username" form:"username"`                            //角色名称
	Status   int `gorm:"column:create_by" db:"create_by" json:"createBy" form:"create_by"`         //创建人
	CreatedAt int `gorm:"column:create_time" db:"create_time" json:"createTime" form:"create_time"` //创建时间
	UpdatedAt int `gorm:"column:update_time" db:"update_time" json:"updateTime" form:"update_time"` //更新时间
}

// 设置表名
func TableName() string {
	return "tpsw_user"
}

// createUser 创建用户
func CreateUser(user User) *gorm.DB {
	result := GetDB().Table(TableName()).Create(&user)
	return result
}

// 查找用户
func FindForId(id int) *User {
	user := User{}
	GetDB().Table(TableName()).Where("uid = ?", id).Find(&user)
	return &user
}

// 更新数据
func UpdateUserForId(id int, data map[string]interface{}) * User {
	user := User{}
	GetDB().Table(TableName()).Model(&user).Where("uid = ?", id).Updates(data)
	return &user
}

// 删除用户
func DeleteUser(id int) *gorm.DB {
	user := User{}
	return GetDB().Table(TableName()).Delete(&user, id)
}