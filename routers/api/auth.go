package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"gin_example/pkg/app"
	"gin_example/pkg/e"
	"gin_example/pkg/logging"
	"gin_example/service/user_service"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func Auth(c *gin.Context) {
	appG := app.Gin{C: c}
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	userService := user_service.User{
		Username: username,
		Password: password,
	}

	if !userService.Auth() {
		appG.Response(http.StatusOK, e.ERROR_AUTH, nil)
		return
	}

	token, err := userService.GetToken()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	data := make(map[string]interface{})
	data["token"] = token
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetInfo(c *gin.Context) {
	// appG := app.Gin{C: c}
	// token := c.Query("token")

}

func InfoForTest(c *gin.Context) {
	data := make(map[string]interface{})
	data["name"] = "非法操作"
	data["avatar"] = "https://cn.vuejs.org/images/logo.png"
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": data,
	})
}
