package models

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model

	Name string `json:"name"`
}

func GetTagsByArticleID(articleID int) ([]*Tag, error) {
	var tags []*Tag
	rows, err := db.Raw("select t.* from blog_tag t inner join blog_article_tag at on t.id = at.tag_id where at.article_id = ?", articleID).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tag Tag
		_ = db.ScanRows(rows, &tag)
		tags = append(tags, &tag)
	}
	return tags, nil
}

func GetTags(pageNum int, pageSize int, maps interface{}) ([]*Tag, error) {
	var tags []*Tag
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func GetTagTotal(maps interface{}) (int, error) {
	count := 0
	err := db.Model(&Tag{}).Where(maps).Count(&count).Error
	return count, err
}

func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ?", name).First(&tag).Error
	if err != nil {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

func AddTag(name string, state int, createdBy string) error {
	tag := Tag{
		Name: name,
	}
	if err := db.Create(&tag).Error; err != nil {
		return err
	}

	return nil
}

func ExistTagByID(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ?", id).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

func DeleteTag(id int) error {
	err := db.Where("id = ?", id).Delete(&Tag{}).Error

	return err
}

func EditTag(id int, data interface{}) error {
	err := db.Model(&Tag{}).Where("id = ?", id).Updates(data).Error

	return err
}
