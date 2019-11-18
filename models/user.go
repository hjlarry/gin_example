package models

type User struct {
	Model
	Email    string `json:"email"`
	Username string `json:"username" gorm:"unique_index"`
	Password string
	Active   bool `json:"active"`
}

type GithubUser struct {
	Model
	Gid     int
	Email   string `json:"email"`
	Name    string `json:"username" gorm:"unique_index"`
	Picture string
	Link    string
}

func CheckAuth(username, password string) bool {
	var user User
	db.Select("id").Where(User{Username: username, Password: password}).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
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
