package database

import (
	"fmt"
	"linkforge-core/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.DataBase.Username,
		config.Conf.DataBase.Password,
		config.Conf.DataBase.Host,
		config.Conf.DataBase.Port,
		config.Conf.DataBase.Dbname,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	fmt.Println("数据库连接成功")
	return nil
}
