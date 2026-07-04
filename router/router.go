package routers

import (
	"fmt"
	"linkforge-core/config"
	"linkforge-core/middleware"
	"linkforge-core/router/auth"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.JWTAuth())

	auth.InitAuthRouter(r)

	r.Run(fmt.Sprintf(":%d", config.Conf.App.Port))
	return r
}
