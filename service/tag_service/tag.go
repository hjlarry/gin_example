package tag_service

import (
	"encoding/json"
	"gin_example/models"
	"gin_example/pkg/gredis"
	"gin_example/pkg/logging"
	"gin_example/service/cache_service"
)

type Tag struct {
	ID         int
	Name       string
	State      int
	CreatedBy  string
	ModifiedBy string

	PageNum  int
	PageSize int
}

func (t *Tag) GetAll() ([]*models.Tag, error) {
	var tags, cacheTags []*models.Tag

	cache := cache_service.Tag{
		ID:       t.ID,
		State:    t.State,
		PageNum:  t.PageNum,
		PageSize: t.PageSize,
	}
	key := cache.GetTagsKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheTags)
			return cacheTags, nil
		}
	}
	tags, err := models.GetTags(t.PageNum, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}
	gredis.Set(key, tags, 3600)
	return tags, nil
}

func (t *Tag) ExistByID() (bool, error) {
	return models.ExistTagByID(t.ID)
}

func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.Name)
}

func (t *Tag) Count() (int, error) {
	return models.GetTagTotal(t.getMaps())
}

func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	if t.State != -1 {
		maps["state"] = t.State
	}

	return maps
}

func (t *Tag) Add() error {
	tag := map[string]interface{}{
		"name":       t.Name,
		"created_by": t.CreatedBy,
		"state":      t.State,
	}
	if err := models.AddTag(tag); err != nil {
		return err
	}
	return nil
}

func (t *Tag) Edit() error {
	tag := map[string]interface{}{
		"name":        t.Name,
		"modified_by": t.ModifiedBy,
		"state":       t.State,
	}
	if err := models.EditTag(t.ID, tag); err != nil {
		return err
	}
	return nil
}

func (t *Tag) Delete() error {
	return models.DeleteTag(t.ID)
}
