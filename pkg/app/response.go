package app

import (
	"linkforge-core/pkg/e"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func (g *Gin) Response(httpCode, businessCode int, data any) {
	g.C.JSON(httpCode, Response{
		businessCode,
		e.GetResponseMsg(businessCode),
		data,
	})
}
