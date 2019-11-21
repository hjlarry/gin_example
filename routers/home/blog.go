package home

import (
	"gin_example/pkg/logging"
	"gin_example/service/article_service"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/unknwon/com"
	"html/template"
	"net/http"
)

func Index(c *gin.Context) {
	articleService := article_service.Article{
		Status:   true,
		PageSize: 10,
	}

	total, err := articleService.Count()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	articles, err := articleService.GetAll()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":    "My Blog",
		"articles": articles,
		"total":    total,
	})
}

func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	articleService := article_service.Article{ID: id}
	article, err := articleService.Get()
	if err != nil {
		logging.Info(err)
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}

	content := markdown.ToHTML([]byte(article.Content), nil, nil)
	htmlContent := template.HTML(content)

	c.HTML(http.StatusOK, "article.html", gin.H{
		"article":     article,
		"htmlContent": htmlContent,
	})
}
