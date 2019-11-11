package models

import (
	"github.com/jinzhu/gorm"

	"time"
)

type Article struct {
	Model

	Title         string `json:"title"`
	AuthorID      int    `json:"title"`
	Slug          string `json:"slug"`
	Summary       string `json:"summary"`
	CanComment    bool   `json:"can_comment"`
	Status        int    `json:"status"`
	Type          int    `json:"type"`
	Content       string `json:"content" gorm:"type:longtext"`
	CoverImageUrl string `json:"cover_image_url"`

	Tags []*Tag `gorm:"-"`
	User User   `gorm:"-"`
}

type ArticleTag struct {
	Model
	PostId int `gorm:"index"`
	TagId  int `gorm:"index"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetArticleTotal(maps interface{}) (int, error) {
	var count int
	err := db.Model(&Article{}).Where(maps).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetArticles(pageNum int, pageSize int, maps interface{}) ([]*Article, error) {
	var articles []*Article
	err := db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}

func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	//err = db.Model(&article).Related(&article.Tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, err
}

func EditArticle(id int, data interface{}) error {
	err := db.Model(&Article{}).Where("id = ?", id).Updates(data).Error

	return err
}

func AddArticle(data map[string]interface{}) error {
	err := db.Create(&Article{
		Title:         data["title"].(string),
		Content:       data["content"].(string),
		CoverImageUrl: data["cover_image_url"].(string),
		//CreatedOn:     data["created_by"].(string),
	}).Error

	return err
}

func DeleteArticle(id int) error {
	err := db.Where("id = ?", id).Delete(Article{}).Error
	return err
}
