package api

import (
	"assemble/config"
	"assemble/logger/model/rest"

	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/v3"
	"github.com/medivhzhan/weapp/v3/phonenumber"
)

var Store = map[string]*weapp.LoginResponse{}

func Login(c *gin.Context) error {
	if c.Query("code") == "" {
		rest.Failed(c, "code 不允许为空")
		return nil
	}

	sdk := weapp.NewClient(config.GetApplets().AppId, config.GetApplets().Secret)
	number, err := sdk.NewPhonenumber().GetPhoneNumber(&phonenumber.GetPhoneNumberRequest{Code: c.Query("code")})
	if err != nil {
		return err
	}

	login, err := sdk.Login(c.Query("code"))
	if err != nil {
		return err
	}

	Store[login.OpenID] = login
	// todo 存储数据库，用户ID 关联
	rest.Success(c, map[string]any{"phone": number, "openId": login.OpenID})
	return nil
}
