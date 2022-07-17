package main

import (
	"framework/bootstrap"
	"framework/router"
)

func main() {
	// 加载环境变量
	config := bootstrap.LoadEnvironmentVariables()

	_, err := bootstrap.Database(config)
	if err != nil {
		panic(err)
	}

	// 启动路由
	r := router.NewHttpServer()

	r.Run()
}
