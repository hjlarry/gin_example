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
