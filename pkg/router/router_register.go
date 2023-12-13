package router

import (
	"assemble/logger/model/rest"

	"github.com/gin-gonic/gin"
)

var _hooks []Hook

type Hook func(Router)

type Handler func(*gin.Context) error

type Router interface {
	Group(string, ...gin.HandlerFunc) *gin.RouterGroup
	GET(string, Handler)
	POST(string, Handler)
	DELETE(string, Handler)
	PATCH(string, Handler)
	PUT(string, Handler)
}

func Register(hook Hook) {
	_hooks = append(_hooks, hook)
}

type myRouter gin.RouterGroup

func (w *myRouter) Group(s string, handler ...gin.HandlerFunc) *gin.RouterGroup {
	return (*gin.RouterGroup)(w).Group(s, handler...)
}

func (w *myRouter) GET(s string, handler Handler) {
	(*gin.RouterGroup)(w).GET(s, w.wrapper(handler))
}

func (w *myRouter) POST(s string, handler Handler) {
	(*gin.RouterGroup)(w).POST(s, w.wrapper(handler))
}

func (w *myRouter) DELETE(s string, handler Handler) {
	(*gin.RouterGroup)(w).DELETE(s, w.wrapper(handler))
}

func (w *myRouter) PATCH(s string, handler Handler) {
	(*gin.RouterGroup)(w).PATCH(s, w.wrapper(handler))
}

func (w *myRouter) PUT(s string, handler Handler) {
	(*gin.RouterGroup)(w).PUT(s, w.wrapper(handler))
}

func (w *myRouter) wrapper(handler Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := handler(c)
		if err == nil {
			return
		}
		switch e := err.(type) {
		case rest.Rest:
			c.JSON(e.Code.HttpStatus(), e)
		case *rest.Rest:
			c.JSON(e.Code.HttpStatus(), e)
		default:
			ee := rest.Wrap(rest.CodeUnknownError, e)
			c.JSON(ee.Code.HttpStatus(), ee)
		}
	}
}
