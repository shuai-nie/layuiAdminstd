package model

import (
	"fmt"
	// 开放式跟踪器
	otgorm "github.com/eddycjy/opentracing-gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"layuiAdminstd/global"
	"layuiAdminstd/pkg/setting"
	"time"
)

const (
	STATE_OPEN = 1
	STATE_CLOST = 0
)

type Model struct {
	ID uint32
	CreatedBy string
	ModifiedBy string
	CreatedOn uint32
	ModifiedOn uint32
	DeletedOn uint32
	IsDel uint8
}

// 初始化数据库链接
func NewDBEngine(databaseSetting *setting.DatabaseSettings) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf(s,
		databaseSetting.Username,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
		))

	if err != nil {
		return nil, err
	}

	fmt.Sprintf("11111111111111111111111111111111111111111")

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	// 三个对应的回调已经写好并替换掉默认的钩子回调
	// 注册 Callbacks
	//db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	otgorm.AddGormCallbacks(db)
	return db, nil
}
/**
结合 GORM 完成了新增、更新、查询的 Callbacks，在实际项目中常常也是这么使用
毕竟，一个钩子的事，就没有必要自己手写过多不必要的代码了
（注意，增加了软删除后，先前的代码需要增加 deleted_on 的判断）
 */
// 注册回调，
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		newTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(newTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(newTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope ) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		// scope.Get("gorm:delete_option") 检查是否手动指定了 delete_option
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		// scope.FieldByName("DeletedOn") 获取我们约定的删除字段，若存在则 UPDATE 软删除，若不存在则 DELETE 硬删除
		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		isDelField, hasIsDelField := scope.FieldByName("IsDel")
		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDelField {
			now := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
				)).Exec()
		}else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				// scope.QuotedTableName() 返回引用的表名，这个方法 GORM 会根据自身逻辑对表名进行一些处理
				scope.QuotedTableName(),
				// scope.CombinedConditionSql() 返回组合好的条件 SQL，看一下方法原型很明了
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}