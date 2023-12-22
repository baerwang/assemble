package jwt

import (
	"assemble/pkg/context"
	"assemble/pkg/service"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

const (
	_identityKey = "key"
)

func NewAuthMiddleware(realm, secretKey string) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       realm,
		Key:         []byte(secretKey),
		Timeout:     30 * time.Minute,
		MaxRefresh:  30 * time.Minute,
		IdentityKey: _identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(map[string]interface{}); ok {
				return jwt.MapClaims{_identityKey: v}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			return jwt.ExtractClaims(c)[_identityKey].(context.Metadata)
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			return service.Login(c)
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			data, ok := data.(context.Metadata)
			if ok {
				c.Set("session", data)
			}
			return ok
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{"code": code, "message": message})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}
