package page

import (
	"assemble/pkg/constant"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page  int
	Limit int
}

func Get(c *gin.Context) Pagination {
	value, _ := c.Get(constant.Pagination)
	return value.(Pagination)
}
