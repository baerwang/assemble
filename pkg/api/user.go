package api

import (
	"assemble/config"
	"assemble/logger/model/rest"
	"assemble/pkg/context"
	"assemble/pkg/dao"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/v3"
	"github.com/medivhzhan/weapp/v3/phonenumber"
)

func UpdateMobile(c *gin.Context) error {
	sdk := weapp.NewClient(config.GetApplets().AppId, config.GetApplets().Secret,
		weapp.WithHttpClient(&http.Client{Timeout: 5 * time.Second}))

	number, err := sdk.NewPhonenumber().GetPhoneNumber(&phonenumber.GetPhoneNumberRequest{Code: c.Query("code")})
	if err != nil {
		return err
	}

	if err = number.GetResponseError(); err != nil {
		return err
	}

	mobile := number.Data
	rsp := context.GetMetadata(c)
	if err = dao.UpdateUserMobile(mobile.PhoneNumber, mobile.PurePhoneNumber, mobile.CountryCode, rsp.OpenId); err != nil {
		return err
	}
	rest.OK(c)
	return nil
}
