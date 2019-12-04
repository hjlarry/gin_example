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
	User *User  `gorm:"-" json:"user"`
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
	//无法将created_at的值通过scanrows放到article的basemodel中的created_at字段，此方案后续再优化
	//var tags []*Tag
	//rows, err := db.Raw("select b3.*, b1.* from blog_article as b1 left join blog_article_tag as b2 on b1.id=b2.article_id left join blog_tag as b3 on b2.tag_id=b3.id where b1.id = ?", id).Rows()
	//if err != nil {
	//	return nil, err
	//}
	//defer rows.Close()
	//for rows.Next() {
	//	var tag Tag
	//	_ = db.ScanRows(rows, &tag)
	//	_ = db.ScanRows(rows, &article)
	//	tags = append(tags, &tag)
	//}
	//fmt.Printf("%v99", rows)
	//article.Tags = tags
	//return &article, nil
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
		AuthorID:   data["author_id"].(int),
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
