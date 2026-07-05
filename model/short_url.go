package model

import (
	"linkforge-core/database"
	"linkforge-core/pkg/e"
	"linkforge-core/pkg/hashUtil"
	"linkforge-core/types/dto"
	"time"

	"gorm.io/gorm"
)

type ShortUrl struct {
	ID          int       `json:"id"`
	LongUrl     string    `json:"long_url"`
	ShortCode   string    `json:"short_code"`
	CreatedTime time.Time `json:"created_time"`
	ExpiredTime time.Time `json:"expired_time"`
	UpdateTime  time.Time `json:"update_time"`
	IsExpired   int       `json:"is_expired"`
	UserId      int       `json:"user_id"`
	ClickCount  int       `json:"click_count"`
}

func (ShortUrl) TableName() string {
	return "short_url"
}

func CreateShortUrlRecord(createShortUrlRecord *dto.CreateShortUrlRecord) (errorCode int, newShortCode string) {
	expiredTime, _ := time.Parse("2006-01-02 15:04:05", createShortUrlRecord.UserInputParams.ExpiredTime)
	shortCode := ""
	newShortUrlRecord := ShortUrl{
		LongUrl:     createShortUrlRecord.UserInputParams.LongUrl,
		CreatedTime: time.Now(),
		ExpiredTime: expiredTime,
		UpdateTime:  time.Now(),
		UserId:      createShortUrlRecord.UserId,
		IsExpired:   0,
		ClickCount:  0,
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		create_res := tx.Model(ShortUrl{}).Create(&newShortUrlRecord)
		if create_res.Error != nil || create_res.RowsAffected == 0 {
			return create_res.Error
		}

		databaseId := newShortUrlRecord.ID
		shortCode = hashUtil.GenerateShortCode(uint64(databaseId), "")

		update_res := tx.Model(ShortUrl{}).Where("id=?", databaseId).Updates(ShortUrl{
			ShortCode: shortCode,
		})

		if update_res.Error != nil || update_res.RowsAffected == 0 {
			return update_res.Error
		}

		return nil
	})
	if err != nil {
		return e.SHORT_URL_ERROR_CREATE_EXECUTE, ""
	}

	return e.SHORT_URL_SUCCESS, shortCode
}
