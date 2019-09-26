package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
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

func AddTag(data map[string]interface{}) error {
	err := db.Create(&Tag{
		Name:      data["name"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	}).Error

	return err
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

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
