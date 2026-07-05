package service

import (
	"linkforge-core/model"
	"linkforge-core/pkg/app"
	"linkforge-core/pkg/e"
	"linkforge-core/types/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateShortUrlRecord(ctx *gin.Context) {
	appG := app.Gin{C: ctx}
	userInputParams := dto.UserInputShortUrlParams{}
	if err := ctx.ShouldBindJSON(&userInputParams); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	userId, ok := ctx.Get("user_id")

	if !ok {
		appG.Response(http.StatusOK, e.INVALID_TOKEN, nil)
		return
	}

	newRecord := dto.CreateShortUrlRecord{
		UserInputParams: dto.UserInputShortUrlParams{},
		UserId:          userId.(int),
	}

	createCode, shortCode := model.CreateShortUrlRecord(&newRecord)
	if createCode != e.SHORT_URL_SUCCESS {
		appG.Response(http.StatusOK, createCode, nil)
		return
	}
	appG.Response(http.StatusOK, e.SHORT_URL_SUCCESS, shortCode)
}
