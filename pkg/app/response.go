package app

import "github.com/gin-gonic/gin"

type Gin struct {
	c *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Gin) Response(code int, msg string, data interface{}) {
	g.c.JSON(code, Response{
		code,
		msg,
		data,
	})
}
