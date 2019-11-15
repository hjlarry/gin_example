package models

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model

	Name string `json:"name"`
}

type ArticleTag struct {
	Model
	ArticleID int `gorm:"index"`
	TagId     int `gorm:"index"`
}

func GetTagsByArticleID(articleID int) (tags []*Tag, err error) {
	rows, err := db.Raw("select t.* from blog_tag t inner join blog_article_tag at on t.id = at.tag_id where at.article_id = ?", articleID).Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var tag Tag
		_ = db.ScanRows(rows, &tag)
		tags = append(tags, &tag)
	}
	return
}

func GetTags(pageNum int, pageSize int, maps interface{}) ([]*Tag, error) {
	var tags []*Tag
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func GetTagIDByName(name string) int {
	var tag Tag
	db.First(&tag, "name=?", name)
	return tag.ID
}

func isInArray(name string, tagNameArray []string) bool {
	for _, v := range tagNameArray {
		if name == v {
			return true
		}
	}
	return false
}

func getTagArray(aTags, bTags []string) (tags []string) {
	for _, v1 := range aTags {
		if !isInArray(v1, bTags) {
			tags = append(tags, v1)
		}
	}
	return
}

func UpdateMultiTags(originTags, newTags []string, article_id int) error {
	needDelTags := getTagArray(originTags, newTags)
	var needToDelTagID []int
	for _, v := range needDelTags {
		tagID := GetTagIDByName(v)
		needToDelTagID = append(needToDelTagID, tagID)
	}
	db.Delete(&ArticleTag{}, "tag_id in ( ? )", needToDelTagID)

	needAddTags := getTagArray(newTags, originTags)
	for _, v := range needAddTags {
		tag := Tag{Name: v}
		db.FirstOrCreate(&tag, "name=?", v)
		db.Create(&ArticleTag{ArticleID: article_id, TagId: tag.ID})
	}

	return nil
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
