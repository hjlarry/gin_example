package user_service

import (
	"gin_example/models"
	"gin_example/pkg/setting"
	"gin_example/pkg/util"
)

func (u *User) Auth() bool {
	user, err := models.GetUserByName(u.Username)
	if err != nil {
		return false
	}
	enc_password := util.EncodeSha1(u.Password + setting.AppSetting.AuthSalt)
	if user.Password != enc_password {
		return false
	}
	return true
}

func (u *User) GetToken() (string, error) {
	return util.GenerateToken(u.Username, u.Password)
}
