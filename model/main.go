package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SmTest struct {
	Id int
	Name string
	Age int
}

func (SmTest) TableName() string {
	return "sm_test"
}

func main3() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/sm_db?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 单条数据插入
	test := SmTest{
		Name : "yzl",
		Age: 18,
	}

	res := db.Create(&test)			// 添加数据
	fmt.Println(test.Id)			// 添加数据的ID
	fmt.Println(res.Error)			// 添加时出现的错误，没有错误返回nil
	fmt.Println(res.RowsAffected)	// 插入记录的条数

	//多条插入
	// tests := []SmTest{
	// 	{Name: "yzl", Age: 20},
	// 	{Name: "yyy", Age: 21},
	// 	{Name: "zzz", Age: 22},
	// 	{Name: "lll", Age: 23},
	// }
	// res := db.Create(&tests)
	// for _, test := range tests {
	// 	fmt.Println(test.Id)
	// }
	// fmt.Println(res.Error)
	// fmt.Println(res.RowsAffected)

	//单条查询
	// test := SmTest{}
	// res := db.First(&test) //取出按id正序排的第一条
	// res := db.Take(&test) //取出不排序的第一条
	// res := db.Last(&test) //取出按id倒序排的第一条
	// fmt.Println(test)
	// fmt.Println(res.RowsAffected) //返回查到的记录数
	// fmt.Println(res.Error)        //返回错误

	//根据主键查询
	// test := SmTest{}
	// db.First(&test, 7) //查询主键为2的记录
	// fmt.Println(test)
	//根据主键多条查询
	// tests := []SmTest{}
	// db.Find(&tests, []int{1, 2, 7}) //查询主键为1，2，3的记录
	// fmt.Println(tests)

	// 查看全部记录
	// tests := []SmTest{}
	// res := db.Find(&tests)
	// fmt.Println(tests)
	// fmt.Println(res.RowsAffected)
	// fmt.Println(res.Error)

	//String条件查询
	// tests := []SmTest{}
	// res := db.Where("name=?", "yzl").Find(&tests) //根据name精确查询
	// res := db.Where("name like ?", "%y%").Find(&tests) //根据name模糊查询
	// res := db.Where("name in ?", []string{"yzl", "zzz"}).Find(&tests) //根据name in查询
	// res := db.Where("id in (?)", []int{1, 2}).Find(&tests) //根据id in查询
	// res := db.Select("id, name, age").Where("name = ? and age = ?", "yzl", 18).Limit(3).Offset(0).Order("id desc").Find(&tests)
	// fmt.Println(tests)
	// fmt.Println(res.RowsAffected)
	// fmt.Println(res.Error)

	// Struct结构体条件查询
	// test := SmTest{}
	// db.Where(&SmTest{Name: "yzl", Age: 18}).First(&test)
	// // SELECT * FROM sm_test WHERE name = "yzl" AND age = 18 ORDER BY id LIMIT 1;
	// fmt.Println(test)

	// Map条件查询
	// tests := []SmTest{}
	// db.Where(map[string]interface{}{"name": "yzl", "age": 18}).Find(&tests)
	// // SELECT * FROM sm_test WHERE name = "yzl" AND age = 18;
	// fmt.Println(tests)

	// 主键切片条件
	// tests := []SmTest{}
	// db.Where([]int64{3, 4, 7}).Find(&tests)
	// fmt.Println(tests)

	//单列更新
	// res := db.Model(&SmTest{}).Where("age = ?", 18).Update("name", "hello")
	// fmt.Println(res.RowsAffected)

	//多列更新
	// res := db.Model(&SmTest{}).Where("age = ?", 18).Updates(SmTest{Name: "test", Age: 23})
	// fmt.Println(res.RowsAffected)

	//删除
	// res := db.Delete(&SmTest{}, 7) //删除数据
	// fmt.Println(res.RowsAffected)

	//根据条件删除
	// res := db.Where("name = ?", "test").Delete(&SmTest{})
	// fmt.Println(res.RowsAffected)

	//gorm2版本以上不需要使用
	// defer db.Close()


}
