package service

import (
	"assemble/config"
	"assemble/pkg/context"
	"assemble/pkg/dao"
	"assemble/pkg/datajson"
	"assemble/pkg/dto"
	"assemble/pkg/model"
	"assemble/pkg/utils"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/v3"
	"github.com/medivhzhan/weapp/v3/auth"
)

func Login(c *gin.Context) (any, error) {
	var userinfo dto.UserInfo
	if err := c.ShouldBindJSON(&userinfo); err != nil {
		return nil, err
	}

	sdk := weapp.NewClient(config.GetApplets().AppId, config.GetApplets().Secret,
		weapp.WithHttpClient(&http.Client{Timeout: 5 * time.Second}))

	auth2 := sdk.NewAuth()
	session, err := auth2.Code2Session(&auth.Code2SessionRequest{Appid: config.GetApplets().AppId,
		Secret: config.GetApplets().Secret, GrantType: "authorization_code", JsCode: userinfo.Code})
	if err != nil {
		return nil, err
	}

	if err = session.GetResponseError(); err != nil {
		return nil, err
	}

	if len(session.Openid) == 0 {
		return nil, errors.New("OpenID 获取失败不允许为空")
	}

	flag, err := dao.FindUserOpenId(session.Openid)
	if err != nil {
		return nil, err
	}

	if !flag {
		info, err := sdk.DecryptUserInfo(session.SessionKey, userinfo.RawData,
			userinfo.EncryptedData, userinfo.Signature, userinfo.IV)
		if err != nil {
			return nil, err
		}

		current := datajson.Datetime(time.Now())
		user := model.User{Id: utils.GoId(), UnionId: session.Unionid, OpenId: session.Openid, Logout: 2,
			Sex: info.Gender, NickName: info.Nickname, Logo: info.Avatar,
			City: info.City, Country: info.Country,
			CreateDate: &current, UpdateDate: &current}

		if err = dao.CreateUser(user); err != nil {
			return nil, err
		}
	}
	return context.Metadata{OpenId: session.Openid, UnionId: session.Unionid, SessionKey: session.SessionKey}, nil
}
