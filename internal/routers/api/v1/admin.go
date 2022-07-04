package v1

import "github.com/gin-gonic/gin"

type Admin struct {}

func NewAdmin() Admin {
	return Admin{}
}

func (a Admin) Get(c *gin.Context) {}
func (a Admin) List(c *gin.Context) {}
func (a Admin) Create(c *gin.Context) {}
func (a Admin) Update(c *gin.Context) {}
func (a Admin) Delete(c *gin.Context) {}
