package logger

import (
	"net"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"assemble/logger/model/rest"

	"github.com/gin-gonic/gin"
)

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		logger.Infof("status:%d resp_time:%v method:%s path:%s query:%s ip:%s ", c.Writer.Status(), time.Since(start),
			c.Request.Method, c.Request.URL.Path, c.Request.URL.RawQuery, c.ClientIP())
	}
}

// ZapRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func ZapRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				if brokenPipe {
					logger.Errorf("error：%s", err)
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Errorf("[Recovery from panic] error：%T stack：%s", err, string(debug.Stack()))
				} else {
					logger.Errorf("[Recovery from panic] error：%T", err)
				}
				c.AbortWithStatusJSON(http.StatusInternalServerError, rest.Rest{Code: http.StatusInternalServerError, Msg: "后台服务异常请联系管理员"})
			}
		}()
		c.Next()
	}
}
