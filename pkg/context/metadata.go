package context

import "github.com/gin-gonic/gin"

type Metadata struct {
	OpenId     string
	UnionId    string
	SessionKey string
}

func GetMetadata(c *gin.Context) Metadata {
	return c.MustGet("session").(Metadata)
}
