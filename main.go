package main

import (
	"fmt"
	"linkforge-core/config"
	"linkforge-core/datebase"
	routers "linkforge-core/router"
	"os"
)

func main() {
	configErr := config.InitConfig()
	if configErr != nil {
		fmt.Println("配置加载失败----")
		os.Exit(1)
		return
	}

	dbErr := datebase.InitDB()
	if dbErr != nil {
		fmt.Println("数据库连接失败")
		os.Exit(1)
		return
	}

	routers.InitRouter()
}
