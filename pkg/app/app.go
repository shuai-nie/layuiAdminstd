package app
// εεΊε€η
import (
	"github.com/gin-gonic/gin"
	"layuiAdminstd/pkg/errcode"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	 Page int
	 PageSize int
	 TotalRows int
}
type data struct {
	code uint8
	data interface{}
	count int
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	/*r.Ctx.JSON(http.StatusOK, gin.H{
		"data" : list,
		"pager" : Pager{
			Page: GetPage(r.Ctx),
			PageSize: GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})*/

	r.Ctx.JSON(http.StatusOK, gin.H{
		"data" : list,
		"count" : totalRows,
		"code" : 0,
	})
	/*r.Ctx.JSON(http.StatusOK, data{
		data: list,
		count: totalRows,
		code: 0,
	})*/
}

func (r *Response) ToErrorResponse(err *errcode.Error ) {
	response := gin.H{"code":err.Code(), "msg":err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}