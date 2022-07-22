package errcode

var(
	ErrorGetAdminListFail = NewError(20010001, "获取列表失败")
	ErrorCreateAdminFail = NewError(20010002, "新建失败")
	ErrorUpdateAdminFail = NewError(20010003, "编辑失败")
	ErrorDeleteAdminFail = NewError(20010004, "删除失败")


	ErrorUploadFileFail = NewError(20030001, "上传文件失败")

	ErrorCreateAuthGroupFail = NewError(20040001, "新建失败")
	ErrorUpdateAuthGroupRequestFaill = NewError(20040002, "编辑失败")


)
