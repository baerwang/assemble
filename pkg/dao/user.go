package dao

import (
	"assemble/config"
	"assemble/pkg/model"
	"time"
)

func CreateUser(u model.User) error {
	return config.DB.Table("users").Create(&u).Error
}

func FindUserOpenId(openid string) (bool, error) {
	var total int64
	return total != 0, config.DB.Table("users").Where("openid = ? and logout = 2", openid).Count(&total).Error
}

func UpdateUserMobile(phone, purePhone, countryCode string, openid string) error {
	const sql = "openid = ? and logout = 2"
	return config.DB.Table("users").Where(sql, openid).
		Updates(map[string]interface{}{
			"phone": phone, "pure_phone": purePhone, "country_code": countryCode, "update_date": time.Now()}).Error
}
