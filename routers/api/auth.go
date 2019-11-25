package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin_example/pkg/app"
	"gin_example/pkg/e"
	"gin_example/service/user_service"
)

type AuthForm struct {
	Username string `form:"username" valid:"Required; MaxSize(20);MinSize(3)"`
	Password string `form:"password" valid:"Required; MaxSize(20);MinSize(3)"`
}

func Auth(c *gin.Context) {
	appG := app.Gin{C: c}

	var form AuthForm
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	userService := user_service.User{
		Username: form.Username,
		Password: form.Password,
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
func LogOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"msg":  "good",
		"data": nil,
	})
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
