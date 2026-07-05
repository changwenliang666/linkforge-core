package middleware

import (
	"fmt"
	"linkforge-core/pkg/app"
	"linkforge-core/pkg/auth"
	"linkforge-core/pkg/e"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appG := app.Gin{C: ctx}
		requestRouter := ctx.FullPath()
		fmt.Println("fullPath", requestRouter)
		if auth.WhitePageContain(requestRouter) { // 白名单接口，不进行token校验
			ctx.Next()
			return
		}

		authHeader := ctx.GetHeader("Authorization")
		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			appG.Response(http.StatusOK, e.INVALID_TOKEN, nil)
			ctx.Abort()
			return
		}

		userInfo, authError := auth.ParseToken(parts[1])

		if authError != nil {
			appG.Response(http.StatusOK, e.INVALID_TOKEN, nil)
			ctx.Abort()
			return
		}

		ctx.Set("user_id", userInfo.UserId)
		ctx.Set("username", userInfo.Username)
		ctx.Next()
	}
}
