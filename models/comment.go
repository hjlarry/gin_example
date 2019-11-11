package models

type Comment struct {
	Model

	GithubID int
	PostID   int
	RefID    int
	Content  string `gorm:"type:longtext"`
}
