package service

import (
	"linkforge-core/model"
	"linkforge-core/pkg/app"
	"linkforge-core/pkg/e"
	"linkforge-core/pkg/hashUtil"
	"linkforge-core/types/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLogin(ctx *gin.Context) {
	appG := app.Gin{C: ctx}
	userRegistryParams := dto.UserRegistryLoginParams{}
	// 参数校验
	if err := ctx.ShouldBindJSON(&userRegistryParams); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	errCode, token := model.LoginUser(&userRegistryParams)
	if errCode != e.AUTH_LOGIN_SUCCESS {
		appG.Response(http.StatusOK, errCode, nil)
		return
	}

	appG.Response(http.StatusOK, errCode, token)

}

func UserRegistry(ctx *gin.Context) {
	appG := app.Gin{C: ctx}
	userRegistryParams := dto.UserRegistryLoginParams{}
	// 参数校验
	if err := ctx.ShouldBindJSON(&userRegistryParams); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	// 加密密码
	encodePassword, err := hashUtil.HashEncode(userRegistryParams.Password)
	if err != nil {
		appG.Response(http.StatusOK, e.AUTH_REGISTRY_ERROR_GENERATE, nil)
		return
	} else {
		userRegistryParams.Password = encodePassword
	}

	errCode := model.RegistryUser(&userRegistryParams)
	if errCode != e.AUTH_REGISTRY_SUCCESS {
		appG.Response(http.StatusOK, errCode, nil)
		return
	}

	appG.Response(http.StatusOK, e.AUTH_REGISTRY_SUCCESS, nil)
}
