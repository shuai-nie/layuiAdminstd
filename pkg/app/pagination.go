package app
// 分页处理
import (
	"github.com/gin-gonic/gin"
	"layuiAdminstd/pkg/convert"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	PageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if PageSize <= 0 {
		// 待修改
		return 10
	}

	if PageSize > 100 {
		return 100
	}
	return PageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		return (page - 1) * pageSize
	}
	return result
}