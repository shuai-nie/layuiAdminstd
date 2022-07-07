# layuiAdminstd
layuiAdminstd-status-template

```
//【增】：向test表插入一个nickname字段，如果该表不存在，则自动建立。
layui.data('test', {  
	key: 'nickname'  ,
	value: 'ABC'
}); 
//【删】：删除test表的nickname字段
layui.data('test', {  
	key: 'nickname'  
	,remove: true
});

layui.data('test', null); 
//删除test表  
//【改】：同【增】，会覆盖已经存储的数据  
//【查】：向test表读取全部的数据
var localTest = layui.data('test');
console.log(localTest.nickname); 
//获得“ABC”

//简单操作
layui.data('cate', {  
	key: 'data'  
	,value: [{    
		key: 'id'    
		,value: 1  
	},{    
		key: 'name'    
		,value: 'abc'  
	}]
});
//取值
var cate = layui.data('cate');
console.log(cate.data);

//复杂操作
layui.data('cate', {        
	key: 'data',        
	value: [            
		{date: 'id', id: 1, content:'00000'}            
		,{date: 'id', id: 2, content:'11111'}            
		,{date: 'id', id: 3, content:'22222'}            
		,{date: 'id', id: 4, content:'33333'}            
	]    
});        
//追加数据    
var cates = layui.data('cate').data;    
cates.push({date: 'id', id: 5, content:'4444444'});        
```
//移除数据    
```cates.splice(2,1);        ```
//更新操作    
```
layui.data('cate', {        
	key: 'data',        
	value: cates    
});        
console.info(layui.data('cate'));
```
/*************************************/
GET请求
两种常见情况

服务器配置
```og.GET("/file/:name", Controller.UploadControl.DownloadFile)```

前端请求URL
```	http://localhost:8082/og/file/test.jpg```

参数获取
```
func (* UploadController)DownloadFile(c *gin.Context)  {
	name := c.Param("name")
}
```

服务器配置
```og.GET("/file", Controller.UploadControl.DownloadFile)```

前端请求URL
```http://localhost:8082/og/file?f=test.jpg```

参数获取
```
func (* UploadController)DownloadFile(c *gin.Context)  {
	name := c.Query("f")
}
```

POST请求
```
data, err := ioutil.ReadAll(c.Request.Body)
CheckError(err)
var msg struct {
	Ids []int
}
json.Unmarshal(data, &msg)
```
FORM请求
用来文件上传

file, _ := c.FormFile("file")

2022年6月20日
导航栏模版二级菜单栏

2022年6月22日
整体规划修改
2022年6月23日
添加URL分组


https://laravelacademy.org/post/21877
https://cba.github.io/layuiAdmin-doc/#login-auth

#/***************************************************/
server {
	listen       80;
	server_name  layui-adminstd-html.nf;
	location / {
		index  index.html index.htm index.php;
		root   "D:/WWW/github.com/layuiAdminstd/layui-template";

		#try_files $uri $uri/ /index.php$is_args$args;
		if (!-e $request_filename){
			rewrite ^/(.*)$ /index.php?s=$1;
		}
		#autoindex  on;
	}
	location /api {
		 proxy_pass                 http://127.0.0.1:8080;
            proxy_redirect             off;
            proxy_set_header           Host             $host;
            proxy_set_header           X-Real-IP        $remote_addr;
            proxy_set_header           X-Forwarded-For  $proxy_add_x_forwarded_for;
	}
}
#/***************************************************/
main.go
````
func main() {
	db := models.DB();
	user := models.User{}
	var userSlice []models.User

	//查询id为1的用户 正序
	//SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = 1 AND ((`users`.`id` = 1)) ORDER BY `users`.`id` ASC LIMIT 1
	db.First(&user, 1)
	//查询id为1的最后一位用户 逆序
	//SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = 2 AND ((`users`.`id` = 1)) ORDER BY `users`.`id` DESC LIMIT 1
	db.Last(&user, 1)

	//Where 条件查询

	//SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((name='hanyun'))
	db.Where("name=?", "hanyun").Find(&userSlice)
	//相等
	//SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((name='hanyun')) ORDER BY `users`.`id` ASC LIMIT 1
	db.Where("name=?", "hanyun").First(&user)

	//不等
	// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((name<>'hanyun'))
	db.Where("name<>?", "hanyun").Find(&userSlice)

	//like
	//SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((name like'%h%'))
	db.Where("name like?", "%h%").Find(&userSlice)

	//and
	//SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((name like'%h%' and gender=1))
	db.Where("name like? and gender=?", "%h%", 1).Find(&userSlice)

	//between and
	//SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((created_at between '2020-03-16 11:35:18' and '2020-03-17 17:40:55'))
	db.Where("created_at between ? and ?", "2020-03-16 11:35:18", "2020-03-17 17:40:55").Find(&userSlice)

	//Struct & Map
	//SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((`users`.`name` = 'hanyun'))
	db.Where(&models.User{Name: "hanyun"}).Find(&userSlice)

	//为什么？当通过结构体进行查询时，GORM将会只通过非零值字段查询，这意味着如果你的字段值为0，''， false 或者其他 零值时，将不会被用于构建查询条件
	//SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL
	db.Where(&models.User{Gender: 0}).Find(&userSlice)
	// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((`users`.`Gender` = 0))
	db.Where(map[string]interface{}{"Gender": 0}).Find(&userSlice)

	// in
	//数值的切片会被当做主键进行查询
	//SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((`users`.`id` IN (8,9)))
	db.Where([]int{8, 9}).Find(&userSlice)

	// not in
	//SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((`users`.`id` NOT IN (8,9)))
	db.Not([]int{8, 9}).Find(&userSlice)

	//or查询
	//SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((`users`.`id` IN (8,9)) OR (`users`.`Gender` = 0))
	db.Where([]int{8, 9}).Or(map[string]interface{}{"Gender": 0}).Find(&userSlice)

	//选择字段
	//SELECT name,gender FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((`users`.`id` IN (8,9)))
	db.Select("name,gender").Where([]int{8, 9}).Find(&userSlice)
	//SELECT name, gender FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((`users`.`id` IN (8,9)))
	db.Select([]string{"name", "gender"}).Where([]int{8, 9}).Find(&userSlice)

	//排序 order
	//单一字段排序
	//SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL ORDER BY name desc
	db.Order("name desc").Find(&userSlice)
	//多字段排序
	//SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL ORDER BY name desc,id asc
	db.Order("name desc,id asc").Find(&userSlice)

	//数量限制
	// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL LIMIT 1
	db.Limit(1).Find(&userSlice)
	//取消数量限制
	//SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL
	db.Limit(1).Limit(-1).Find(&userSlice)

	//偏移
	//SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL LIMIT 1 OFFSET 2
	db.Limit(1).Offset(2).Find(&userSlice)
	//取消偏移
	//SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL LIMIT 1
	db.Limit(1).Offset(2).Offset(-1).Find(&userSlice)

	//总数
	count := 0
	//SELECT count(*) FROM `users`  WHERE `users`.`deleted_at` IS NULL
	db.Find(&userSlice).Count(&count)

	//查询的综合运用，包含了分页
	//SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL LIMIT 1 OFFSET 1
	//SELECT count(*) FROM `users`  WHERE `users`.`deleted_at` IS NULL
	db.Offset(1).Limit(1).Find(&userSlice).Limit(-1).Offset(-1).Count(&count)

	//group
	//SELECT name FROM `users`  WHERE `users`.`deleted_at` IS NULL GROUP BY name
	db.Select("name").Group("name").Find(&userSlice)
	//Having
	//SELECT count(name) c FROM `users`  WHERE `users`.`deleted_at` IS NULL GROUP BY name HAVING (c>1)
	db.Select("count(name) c").Group("name").Having("c>?", 1).Find(&userSlice)
}
````
model.go
````
package models

import (
	"fmt"
	"ginLearn.com/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt utils.JSONTime
	UpdatedAt utils.JSONTime
	DeletedAt *utils.JSONTime `sql:"index"`
}

var db *gorm.DB

func init() {
	setup()
}

// 获得MySQL的资源链接
func DB() *gorm.DB {
	return db
}

// Setup initializes the database instance
func setup() {
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"12345678",
		"127.0.0.1",
		3306,
		"gorm"))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	//采用复数的表名
	db.SingularTable(false)
	//自动数据迁移
	db.AutoMigrate(User{})
	//打印日志
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}
````


````
————configs：配置文件
————global：全忽悠变量
————internal：内部模块
	————dao：数据访问层
	————middleware：http中间件
	————model：模型层，用于存放 model 对象
	————routers：路由相关的逻辑
	————service：项目核心业务逻辑
————pkg：项目相关模块包
————storage：项目生成的临时文件
————scripts：各类构建、安装、分析等操作的脚本
````