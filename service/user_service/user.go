package user_service

import (
	"gin_example/models"
	"gin_example/pkg/setting"
	"gin_example/pkg/util"
)

type User struct {
	ID int

	Username string
	Email    string
	Password string
	Active   bool
}

func (u *User) Add() error {
	enc_password := util.EncodeSha1(u.Password + setting.AppSetting.AuthSalt)
	data := map[string]interface{}{
		"username": u.Username,
		"password": enc_password,
		"email":    u.Email,
		"active":   u.Active,
	}
	return models.AddUser(data)
}
