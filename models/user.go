package models

import "github.com/jinzhu/gorm"

type User struct {
	Model
	Email    string `json:"email"`
	Username string `json:"username" gorm:"unique_index"`
	Password string `json:"-"`
	Active   bool   `json:"active"`
}

type GithubUser struct {
	Model
	Gid     int
	Email   string `json:"email"`
	Name    string `json:"username" gorm:"unique_index"`
	Picture string
	Link    string
}

func AddUser(data map[string]interface{}) error {
	user := User{
		Username: data["username"].(string),
		Email:    data["email"].(string),
		Password: data["password"].(string),
		Active:   data["active"].(bool),
	}
	err := db.Create(&user).Error
	return err
}

func GetUsers(pageNum int, pageSize int, maps map[string]interface{}) ([]*User, error) {
	var users []*User
	maps["deleted_on"] = 0
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Order("id desc").Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return users, nil
}

func GetUserTotal(maps map[string]interface{}) (int, error) {
	var count int
	maps["deleted_on"] = 0
	err := db.Model(&User{}).Where(maps).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetUser(id int) (*User, error) {
	var user User
	err := db.Where("id = ? AND deleted_on = ?", id, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, err
}

func GetUserByName(username string) (*User, error) {
	var user User
	err := db.Where("username = ? AND deleted_on = ?", username, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, err
}

func EditUser(id int, data interface{}) error {
	err := db.Model(&User{}).Where("id = ?", id).Updates(data).Error
	return err
}

func DeleteUser(id int) error {
	err := db.Where("id = ?", id).Delete(User{}).Error
	return err
}

func ExistUserByName(username string) (bool, error) {
	var user User
	err := db.Select("id").Where("username = ?", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}
