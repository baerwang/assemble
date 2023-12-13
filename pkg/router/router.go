package router

import (
	"assemble/pkg/api"
	"assemble/pkg/constant"
	"assemble/pkg/dto"
	"assemble/pkg/model/page"
	"assemble/pkg/router/val"
	"fmt"
	"strconv"

	"assemble/logger"
	"assemble/logger/model/rest"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	router := gin.New()
	router.Use(logger.GinLogger(), logger.ZapRecovery(true))

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("VaDate", val.Date)
	}

	e := router.Group("")
	for _, hook := range _hooks {
		hook((*myRouter)(e))
	}

	return router
}

func checkStd() gin.HandlerFunc {
	return func(c *gin.Context) {
		rsp, ok := api.Store[c.GetHeader("openid")]
		if !ok {
			rest.Failed(c, "当前用户没有登陆")
			c.Abort()
			return
		}
		c.Set("session", rsp)
		c.Next()
	}
}

func checkPageAndIndex(limit int) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Query(constant.Page) == "" || c.Query(constant.Limit) == "" {
			rest.Failed(c, "page or index param invalid")
			c.Abort()
			return
		}

		value, err := strconv.Atoi(c.Query(constant.Page))
		if err != nil || value < 1 {
			rest.Failed(c, "page param invalid")
			c.Abort()
			return
		}
		c.Set(constant.Page, value)

		value, err = strconv.Atoi(c.Query(constant.Limit))
		if err != nil || value < 1 {
			rest.Failed(c, "limit param invalid")
			c.Abort()
			return
		}
		if value > limit {
			rest.Failed(c, fmt.Sprint("limit param limit ", limit))
			c.Abort()
			return
		}

		if c.GetInt(constant.Page) > 0 && value > 0 {
			c.Set(constant.Page, (c.GetInt(constant.Page)-1)*value)
		}

		c.Set(constant.Limit, value)
		c.Set(constant.Pagination, page.Pagination{Page: c.GetInt(constant.Page), Limit: value})

		c.Next()
	}
}

func paramParseIds() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dtoIds dto.Ids
		if err := c.BindJSON(&dtoIds); err != nil {
			rest.Failed(c, err.Error())
			c.Abort()
			return
		}

		ids := make([]int64, 0, len(dtoIds.Ids))

		for _, id := range dtoIds.Ids {
			i, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				rest.Failed(c, err.Error())
				c.Abort()
				return
			}
			ids = append(ids, i)
		}

		c.Set("ids", ids)
	}
}

func paramParseId() gin.HandlerFunc {
	return func(c *gin.Context) {
		parseInt, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			rest.Failed(c, "id param invalid")
			c.Abort()
			return
		}
		c.Set("id", parseInt)
		c.Next()
	}
}

func checkAssert() gin.HandlerFunc {
	return func(c *gin.Context) {
		status, err := strconv.ParseInt(c.Query("status"), 10, 64)
		if err != nil || status < 1 || status > 2 {
			rest.Failed(c, "status param invalid")
			c.Abort()
			return
		}
		c.Set("status", status)
		c.Next()
	}
}
