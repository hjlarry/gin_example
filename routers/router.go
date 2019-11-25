package routers

import (
	"gin_example/middleware/jwt"
	"html/template"
	"net/http"

	_ "gin_example/docs"
	"gin_example/pkg/export"
	"gin_example/pkg/qrcode"
	"gin_example/pkg/setting"
	"gin_example/pkg/upload"
	"gin_example/pkg/util"
	"gin_example/routers/api"
	v1 "gin_example/routers/api/v1"
	"gin_example/routers/home"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.SetFuncMap(template.FuncMap{
		"dateFormat": util.DateFormat,
	})
	r.LoadHTMLGlob("templates/*")
	gin.SetMode(setting.ServerSetting.RunMode)

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))
	r.StaticFS("/static", http.Dir("runtime/static"))

	url := ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.GET("/", home.Index)
	r.GET("/articles/:id", home.GetArticle)

	r.POST("/upload", api.UploadImage)
	r.POST("/api/v1/user/login", api.Auth)
	r.POST("/api/v1/user/logout", api.LogOut)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/user/info", api.GetInfo)

		apiv1.GET("/tags", v1.GetTags)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		r.POST("/tags/export", v1.ExportTag)
		r.POST("/tags/import", v1.ImportTag)

		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)

		apiv1.GET("/users", v1.GetUsers)
		apiv1.POST("/users", v1.AddUser)
		apiv1.GET("/users/:id", v1.GetUser)
		apiv1.PUT("/users/:id", v1.EditUser)
		apiv1.DELETE("/users/:id", v1.DeleteUser)
	}
	return r
}
