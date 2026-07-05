package dto

type UserInputShortUrlParams struct {
	LongUrl     string `json:"long_url" binding:"required,max=1024"`
	ExpiredTime string `json:"expired_time" binding:"required"`
}

type CreateShortUrlRecord struct {
	UserInputParams UserInputShortUrlParams
	UserId          int `json:"user_id"`
}
