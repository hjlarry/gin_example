package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model

	Title         string `json:"title"`
	AuthorID      int    `json:"-"`
	Slug          string `json:"slug"`
	Summary       string `json:"summary"`
	CanComment    bool   `json:"can_comment"`
	Status        bool   `json:"status"`
	Type          int    `json:"-"`
	Content       string `json:"content" gorm:"type:longtext"`
	CoverImageUrl string `json:"cover_image_url"`

	Tags []*Tag `gorm:"-" json:"tags"`
	User User   `gorm:"-" json:"-"`
}

func GetArticleTotal(maps map[string]interface{}) (int, error) {
	var count int
	maps["deleted_on"] = 0
	err := db.Model(&Article{}).Where(maps).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetArticles(pageNum int, pageSize int, maps map[string]interface{}) ([]*Article, error) {
	var articles []*Article
	maps["deleted_on"] = 0
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Order("id desc").Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}

func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ? AND deleted_on = ?", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &article, err
}

func EditArticle(id int, data interface{}) error {
	err := db.Model(&Article{}).Where("id = ?", id).Updates(data).Error
	return err
}

func AddArticle(data map[string]interface{}) (int, error) {

	article := Article{
		Title:      data["title"].(string),
		Slug:       data["slug"].(string),
		Summary:    data["summary"].(string),
		Content:    data["content"].(string),
		CanComment: data["can_comment"].(bool),
		Status:     data["status"].(bool),
		Model: Model{
			CreatedOn: data["created_at"].(*time.Time),
		},
	}

	err := db.Create(&article).Error

	return article.ID, err
}

func DeleteArticle(id int) error {
	err := db.Where("id = ?", id).Delete(Article{}).Error
	return err
}
