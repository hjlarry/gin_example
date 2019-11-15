package article_service

import (
	"encoding/json"
	"gin_example/models"
	"gin_example/pkg/gredis"
	"gin_example/pkg/logging"
	"gin_example/pkg/util"
	"gin_example/service/cache_service"
	"time"
)

type Article struct {
	ID       int
	TagID    int
	Status   int
	PageNum  int
	PageSize int

	Title      string
	Slug       string
	Summary    string
	Content    string
	CanComment bool
	CreatedAt  *time.Time
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
	article.CreatedAt = util.DateFormat(*article.CreatedOn, "2006-01-02 15:04")
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
		Status:   a.Status,
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
	// TODO better query
	for _, a := range articles {
		a.CreatedAt = util.DateFormat(*a.CreatedOn, "2006-01-02 15:04")
		a.Tags, _ = models.GetTagsByArticleID(a.ID)
	}
	_ = gredis.Set(key, articles, 3600)
	return articles, nil
}

func (a *Article) Count() (int, error) {
	return models.GetArticleTotal(a.getMaps())
}

func (a *Article) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	if a.Status != -1 {
		maps["status"] = a.Status
	}
	return maps
}

func (a *Article) Add() (int, error) {
	article := map[string]interface{}{
		"title":       a.Title,
		"slug":        a.Slug,
		"summary":     a.Summary,
		"content":     a.Content,
		"can_comment": a.CanComment,
		"status":      a.Status,
		"created_at":  a.CreatedAt,
	}
	articleId, err := models.AddArticle(article)
	if err != nil {
		return -1, err
	}
	return articleId, nil
}

func (a *Article) Edit() error {
	article := map[string]interface{}{
		"title":  a.Title,
		"status": a.Status,
	}
	if err := models.EditArticle(a.ID, article); err != nil {
		return err
	}
	return nil
}

func (a *Article) Delete() error {
	return models.DeleteArticle(a.ID)
}
