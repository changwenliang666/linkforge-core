package authRouter

import (
	"linkforge-core/service"

	"github.com/gin-gonic/gin"
)

func InitAuthRouter(api *gin.Engine) {
	authGroup := api.Group("/auth")
	{
		authGroup.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, map[string]interface{}{
				"msg": "hello world",
			})
		})

		// 注册接口
		authGroup.POST("/registry", service.UserRegistry)
		// 登录接口
		authGroup.POST("/login", service.UserLogin)
	}
}
