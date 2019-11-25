package v1

import (
	"gin_example/pkg/app"
	"gin_example/pkg/e"
	"gin_example/pkg/util"
	"gin_example/service/user_service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func GetUser(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	userService := user_service.User{ID: id}
	user, err := userService.Get()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_USER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, user)
}

func GetUsers(c *gin.Context) {
	appG := app.Gin{C: c}

	page, limit := util.GetPageAndLimit(c)
	userService := user_service.User{
		PageNum:  page,
		PageSize: limit,
	}

	total, err := userService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_USER_FAIL, nil)
		return
	}

	users, err := userService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_USERS_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = users
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)

}

type AddUserForm struct {
	Username string `form:"username" valid:"Required;MaxSize(20);MinSize(3)"`
	Password string `form:"password" valid:"Required;MaxSize(20);MinSize(3)"`
	Email    string `form:"email" valid:"Email"`
	Active   bool   `form:"active"`
}

func AddUser(c *gin.Context) {
	appG := app.Gin{C: c}
	var form AddUserForm

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	userService := user_service.User{
		Username: form.Username,
		Password: form.Password,
		Email:    form.Email,
		Active:   form.Active,
	}

	exists, err := userService.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_USER_FAIL, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_USER, nil)
		return
	}

	err = userService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

type EditUserForm struct {
	Password string `form:"password"`
	Email    string `form:"email" valid:"Email"`
	Active   bool   `form:"active"`
}

func EditUser(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	var form EditUserForm
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	password := form.Password
	if password != "" {
		valid.MaxSize(password, 20, "password").Message("密码必须少于20位")
		valid.MinSize(password, 3, "password").Message("密码必须大于3位")
	}
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	userService := user_service.User{
		ID:       id,
		Password: password,
		Email:    form.Email,
		Active:   form.Active,
	}
	err := userService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_USER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteUser(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	userService := user_service.User{ID: id}
	err := userService.Delete()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_DELETE_USER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
