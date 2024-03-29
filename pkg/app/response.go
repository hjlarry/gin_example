package app

import (
	"github.com/gin-gonic/gin"

	"gin_example/pkg/e"
)

type Gin struct {
	C               *gin.Context
	CurrentUserName string
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}
