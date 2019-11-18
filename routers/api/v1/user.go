package v1

import (
	"gin_example/pkg/app"
	"gin_example/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	appG := app.Gin{C: c}

	data := make(map[string]interface{})
	data["lists"] = []string{}
	data["total"] = 100

	appG.Response(http.StatusOK, e.SUCCESS, data)

}
