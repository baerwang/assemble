package router

import "assemble/pkg/api"

func init() {
	Register(func(router Router) {
		router.PUT("/api/v1/mobile", api.UpdateMobile)
	})
}
