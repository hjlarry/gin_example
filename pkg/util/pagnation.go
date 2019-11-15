package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"gin_example/pkg/setting"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}

func GetLimit(c *gin.Context) int {
	result := setting.AppSetting.PageSize
	limit, _ := com.StrTo(c.Query("limit")).Int()
	if limit > 0 {
		result = limit
	}
	return result
}

func GetPageAndLimit(c *gin.Context) (page int, limit int) {
	limit = setting.AppSetting.PageSize
	query_limit, _ := com.StrTo(c.Query("limit")).Int()
	if query_limit > 0 {
		limit = query_limit
	}

	page = 0
	query_page, _ := com.StrTo(c.Query("page")).Int()
	if query_page > 0 {
		page = (query_page - 1) * limit
	}
	return
}
