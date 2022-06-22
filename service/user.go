package service

type UserService struct {
}
// 相当于声明一个命名空间
var UserServer UserService
// 查询
func (t *UserService) RoleList(params string, pageNo int, pageSize int) ([]Model.UserRole, int) {
	var total int
	db := GetDb().Model(&Model.UserRole{}).Where(params).Count(&total).Limit(pageSize).Offset((pageNo - 1) * pageSize)
	var list []Model.UserRole
	db = db.Find(&list)
	if db.Error != nil {
		fmt.Println(db.Error)
	}
	return list, total
}
// 新增、编辑
func (t *UserService) RoleEdit(model Model.UserRole, isEdit bool) error {
	db := GetDb()
	if isEdit {
		db = db.Save(model)
	} else {
		// 如果id是自增的，务必带&
		db = db.Create(&model)
	}
	if db.Error != nil {
		return db.Error
	}
	return nil
}
// 删除
func (t *UserService) RoleDel(ids []int) error {
	db := GetDb().Model(&Model.UserRole{})
	db = db.Delete(Model.UserRole{}, ids)
	if db.Error != nil {
		return db.Error
	}
	return nil
}