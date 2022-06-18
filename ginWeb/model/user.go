package model

type UserRole struct {
	Id         int    `gorm:"column:id" db:"id" json:"id" form:"id"`
	Keyword    string `gorm:"column:keyword" db:"keyword" json:"keyword" form:"keyword"`                //角色关键字
	Name       string `gorm:"column:name" db:"name" json:"name" form:"name"`                            //角色名称
	Level      int    `gorm:"column:level" db:"level" json:"level" form:"level"`                        //角色等级
	CreateBy   string `gorm:"column:create_by" db:"create_by" json:"createBy" form:"create_by"`         //创建人
	CreateTime string `gorm:"column:create_time" db:"create_time" json:"createTime" form:"create_time"` //创建时间
	UpdateBy   string `gorm:"column:update_by" db:"update_by" json:"updateBy" form:"update_by"`         //更新人
	UpdateTime string `gorm:"column:update_time" db:"update_time" json:"updateTime" form:"update_time"` //更新时间
}