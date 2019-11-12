package article_service

import (
	"encoding/json"
	"gin_example/models"
	"gin_example/pkg/gredis"
	"gin_example/pkg/logging"
	"gin_example/pkg/util"
	"gin_example/service/cache_service"
)

type Article struct {
	ID       int
	TagID    int
	State    int
	PageNum  int
	PageSize int

	Title string
}

func (a *Article) Get() (*models.Article, error) {
	var cacheArticle *models.Article

	cache := cache_service.Article{ID: a.ID}
	key := cache.GetArticleKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			_ = json.Unmarshal(data, &cacheArticle)
			return cacheArticle, nil
		}
	}
	article, err := models.GetArticle(a.ID)
	tags, err := models.GetTagsByArticleID(a.ID)
	article.Tags = tags
	if err != nil {
		return nil, err
	}
	_ = gredis.Set(key, article, 3600)
	return article, nil
}

func (a *Article) GetAll() ([]*models.Article, error) {
	var articles, cacheArticles []*models.Article

	cache := cache_service.Article{
		TagID:    a.TagID,
		State:    a.State,
		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}
	key := cache.GetArticlesKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			_ = json.Unmarshal(data, &cacheArticles)
			return cacheArticles, nil
		}
	}
	articles, err := models.GetArticles(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}
	for _, a := range articles {
		a.CreatedAt = util.DateFormat(a.CreatedOn, "2006-01-02 15:04")
	}
	_ = gredis.Set(key, articles, 3600)
	return articles, nil
}

func (a *Article) ExistByID() (bool, error) {
	return models.ExistArticleByID(a.ID)
}

func (a *Article) Count() (int, error) {
	return models.GetArticleTotal(a.getMaps())
}

func (a *Article) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	if a.State != -1 {
		maps["state"] = a.State
	}
	return maps
}

func (a *Article) Add() error {
	article := map[string]interface{}{
		"tag_id": a.TagID,
		"title":  a.Title,
		"state":  a.State,
	}
	if err := models.AddArticle(article); err != nil {
		return err
	}
	return nil
}

func (a *Article) Edit() error {
	article := map[string]interface{}{
		"tag_id": a.TagID,
		"title":  a.Title,
		"state":  a.State,
	}
	if err := models.EditArticle(a.ID, article); err != nil {
		return err
	}
	return nil
}

func (a *Article) Delete() error {
	return models.DeleteArticle(a.ID)
}
