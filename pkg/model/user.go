package model

import "assemble/pkg/datajson"

type User struct {
	Id           int64              `json:"id,string,omitempty"`
	NickName     string             `json:"nick_name,omitempty"`                                  // 昵称
	WxNickName   string             `json:"wx_nick_name,omitempty"`                               // 微信昵称
	Sex          int                `json:"sex,omitempty" binding:"oneof=1 2 3"`                  // 性别(1:男 2:女 3:无)
	Name         string             `json:"name,omitempty" binding:"min=1,max=20"`                // 姓名
	Phone        string             `json:"phone,omitempty"`                                      // 有区号手机号 (没有区号手机号可为空, RSA 加密算法存储)
	PurePhone    string             `json:"pure_phone,omitempty"`                                 // 没有区号手机号 (有区号手机号可为空, 加密算法存储)
	CountryCode  string             `json:"country_code,omitempty"`                               // 区号 (没有区号手机号可为空)
	IdentityCard string             `json:"identity_card,omitempty"`                              // 身份证 (身份证 RSA 加密算法存储 - 提现存储)
	Country      string             `json:"country,omitempty" binding:"min=1,max=20"`             // 用户所在国家
	City         string             `json:"city,omitempty" binding:"min=1,max=20"`                // 用户所在城市
	OpenId       string             `json:"openid,omitempty" gorm:"column:openid"`                // 不同小程序不同用户ID
	UnionId      string             `json:"unionid,omitempty" gorm:"column:unionid"`              // 微信开发平台下应用统一用户ID
	Logo         string             `json:"logo,omitempty" binding:"min=1,max=100"`               // 头像链接
	Logout       int                `json:"logout,omitempty"`                                     // 注销账号 (1:是 2:否)
	CreateDate   *datajson.Datetime `json:"create_date,omitempty" gorm:"type:timestamp;not null"` // 创建时间
	UpdateDate   *datajson.Datetime `json:"update_date,omitempty" gorm:"type:timestamp;not null"` // 更新时间
}
