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