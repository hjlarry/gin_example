package models

type User struct {
	Model
	Email    string `json:"email"`
	Name     string `json:"username" gorm:"unique_index"`
	Password string
	Active   int `gorm:"default:'1'"`
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
	db.Select("id").Where(User{Name: username, Password: password}).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}
