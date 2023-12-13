package dto

import "github.com/gin-gonic/gin"

type Ids struct {
	Ids []string `json:"ids"`
}

func GetInt64Array(c *gin.Context) []int64 {
	return c.MustGet("ids").([]int64)
}
