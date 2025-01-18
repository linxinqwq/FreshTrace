package main

import (
	"backed/api/router"
	"backed/app/core/variable"
	_ "backed/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {
	// 引入gin引擎
	engine := gin.Default()

	// 加载路由
	router.MainRouter(engine)

	// 监听路由
	engine.Run(":" + variable.Viper.GetString("application_port"))
}
