package user_service

import (
	"gin_example/models"
	"gin_example/pkg/setting"
	"gin_example/pkg/util"
)

type User struct {
	ID       int
	PageNum  int
	PageSize int
	Token    string

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

func (u *User) GetAll() ([]*models.User, error) {
	users, err := models.GetUsers(u.PageNum, u.PageSize, map[string]interface{}{})
	for _, u := range users {
		u.CreatedAt = util.DateFormat(*u.CreatedOn, "2006-01-02 15:04")
		u.ModifiedAt = util.DateFormat(*u.ModifiedOn, "2006-01-02 15:04")
	}
	return users, err
}

func (u *User) Count() (int, error) {
	return models.GetUserTotal(map[string]interface{}{})
}

func (u *User) Get() (*models.User, error) {
	return models.GetUser(u.ID)
}

func (u *User) GetByName() (*models.User, error) {
	return models.GetUserByName(u.Username)
}

func (u *User) Edit() error {
	data := map[string]interface{}{
		"email":  u.Email,
		"active": u.Active,
	}
	if u.Password != "" {
		data["password"] = util.EncodeSha1(u.Password + setting.AppSetting.AuthSalt)
	}
	return models.EditUser(u.ID, data)
}

func (u *User) Delete() error {
	return models.DeleteUser(u.ID)
}

func (u *User) ExistByName() (bool, error) {
	return models.ExistUserByName(u.Username)
}
