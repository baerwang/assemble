package router

import (
	"assemble/pkg/api"
)

func init() {
	Register(func(router Router) {
		router.GET("/login", api.Login)
	})
}
