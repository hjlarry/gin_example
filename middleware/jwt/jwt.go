package jwt

import (
	"gin_example/pkg/app"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"gin_example/pkg/e"
	"gin_example/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		token := c.Request.Header.Get("X-Token")
		if token == "" {
			appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
			appG.C.Abort()
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			appG.Response(http.StatusOK, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
			return
		}
		if time.Now().Unix() > claims.ExpiresAt {
			appG.Response(http.StatusOK, e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT, nil)
			return
		}
		appG.C.Set("current_user", claims.Username)
		appG.C.Next()
	}
}
