package model

import (
	"errors"
	"linkforge-core/database"
	"linkforge-core/pkg/auth"
	"linkforge-core/pkg/e"
	"linkforge-core/pkg/hashUtil"
	"linkforge-core/types/dto"
	"time"

	"gorm.io/gorm"
)

type LinkUser struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	CreatedTime time.Time `json:"created_time"`
	UpdateTime  time.Time `json:"update_time"`
}

func (LinkUser) TableName() string {
	return "link_user"
}

func RegistryUser(userInfo *dto.UserRegistryLoginParams) (errorCode int) {
	newUser := LinkUser{}
	query_res := database.DB.Model(LinkUser{}).Where("username=?", userInfo.Username).First(&newUser)
	if query_res.Error == nil {
		return e.AUTH_REGISTRY_ERROR_USER_EXIST
	} else if errors.Is(query_res.Error, gorm.ErrRecordNotFound) {
		newUser.Username = userInfo.Username
		newUser.Password = userInfo.Password
		newUser.CreatedTime = time.Now()
		newUser.UpdateTime = time.Now()
		create_res := database.DB.Model(LinkUser{}).Create(&newUser)
		if create_res.Error != nil || create_res.RowsAffected == 0 {
			return e.AUTH_REGISTRY_ERROR_MODEL_EXECUTE
		}
	} else {
		return e.AUTH_REGISTRY_ERROR_MODEL_EXECUTE
	}

	return e.AUTH_REGISTRY_SUCCESS
}

func LoginUser(userInfo *dto.UserRegistryLoginParams) (errorCode int, token string) {
	user := LinkUser{}
	queryRes := database.DB.Model(LinkUser{}).
		Where("username = ?", userInfo.Username).
		First(&user)

	if errors.Is(queryRes.Error, gorm.ErrRecordNotFound) {
		return e.AUTH_LOGIN_ERROR_USER_NOT_EXIST, ""
	}

	if queryRes.Error != nil {
		return e.AUTH_LOGIN_ERROR__MODEL_EXECUTE, ""
	}

	if !hashUtil.CheckHashDecode(userInfo.Password, user.Password) {
		return e.AUTH_LOGIN_ERROR_USER_PASSWORD_ERROR, ""
	}

	tokenStr, err := auth.GenerateToken(user.ID, user.Username)
	if err != nil {
		return e.AUTH_LOGIN_ERROR_GENERATE, ""
	}

	return e.AUTH_LOGIN_SUCCESS, tokenStr
}
