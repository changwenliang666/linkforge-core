package auth

import "github.com/gin-gonic/gin"

func InitAuthRouter(api *gin.Engine) {
	authGroup := api.Group("/auth")
	{
		authGroup.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, map[string]interface{}{
				"msg": "hello world",
			})
		})
		authGroup.POST("/registry")
	}
}
