package routers

import (
	"fmt"
	"linkforge-core/config"
	"linkforge-core/middleware"
	authRouter "linkforge-core/router/auth"
	shortUrlRouter "linkforge-core/router/short_url"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.JWTAuth())

	authRouter.InitAuthRouter(r)
	shortUrlRouter.InitShortUrlRouter(r)

	r.Run(fmt.Sprintf(":%d", config.Conf.App.Port))
	return r
}
