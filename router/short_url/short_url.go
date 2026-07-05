package shortUrlRouter

import (
	"linkforge-core/service"

	"github.com/gin-gonic/gin"
)

func InitShortUrlRouter(api *gin.Engine) {
	shortUrlGroup := api.Group("/api/v1")
	{
		shortUrlGroup.POST("/createShortUrlRecord", service.CreateShortUrlRecord)
	}

}
