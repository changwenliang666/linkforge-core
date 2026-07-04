package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type App struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
}

type Log struct {
	Path       string `mapstructure:"path"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
}

type DataBase struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Dbname   string `mapstructure:"dbname"`
}

type Redis struct {
	ADDR     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type Config struct {
	App      App      `mapstructure:"app"`
	Log      Log      `mapstructure:"log"`
	DataBase DataBase `mapstructure:"database"`
	Redis    Redis    `mapstructure:"redis"`
}

var Conf *Config

func InitConfig() error {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	fmt.Println("current run env:", env)

	v := viper.New()

	v.AddConfigPath("./config")
	v.SetConfigName("config.base")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if env == "dev" {
		v.SetConfigName("config.dev")
		if err := v.MergeInConfig(); err != nil {
			return err
		}
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return err
	}
	Conf = &c
	return nil
}

// // 启用环境变量覆盖（生产环境使用）
// myViper.AutomaticEnv()

// // 手动绑定（docker使用）
// _ = viper.BindEnv("Database.Username", "DB_USERNAME")
// _ = viper.BindEnv("Database.Password", "DB_PASSWORD")
// _ = viper.BindEnv("Database.Host", "DB_HOST")
// _ = viper.BindEnv("Database.Port", "DB_PORT")
// _ = viper.BindEnv("Database.DBName", "DB_NAME")

// _ = viper.BindEnv("Redis.Addr", "REDIS_ADDR")
// _ = viper.BindEnv("Redis.Password", "REDIS_PASSWORD")
// _ = viper.BindEnv("Redis.DB", "REDIS_DB")

// _ = viper.BindEnv("Log.Path", "LOG_PATH")
