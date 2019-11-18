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
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "My Blog",
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
