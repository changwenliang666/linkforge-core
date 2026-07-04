package middleware

import (
	"fmt"
	"linkforge-core/pkg/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestRouter := ctx.FullPath()
		fmt.Println("fullPath", requestRouter)
		if auth.WhitePageContain(requestRouter) { // 白名单接口，不进行token校验
			ctx.Next()
			return
		}

		authHeader := ctx.GetHeader("Authorization")
		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			fmt.Println("token格式不合法")
			ctx.Abort()
			return
		}

		userInfo, authError := auth.ParseToken(parts[1])

		if authError != nil {
			fmt.Println("身份信息认证失败")
			ctx.Abort()
			return
		}

		ctx.Set("userId", userInfo.UserId)
		ctx.Set("username", userInfo.Username)
		ctx.Next()
	}
}
