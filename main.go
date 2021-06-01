package main

import (
	"github.com/gin-gonic/gin"
	"goin/conf"
	"goin/db"
	"goin/logger"
	"goin/middleware"
	"goin/routers"
	"log"
)

// @title simple
// @version 1.0
// @description simple
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @contact.name simple
// @contact.url localhost:8099
// @contact.email 344080699@qq.com
// @host localhost:8099
func main() {
	// 初始化配置文件
	conf.Init()
	err := conf.LoadConf(conf.GetConfPath())
	if err != nil {
		log.Fatalln("load config error", err)
	}

	// 取配置
	config := conf.GetConf()

	// 初始化Mysql数据库
	err = db.MysqlDial(&config.Mysql)
	defer db.EloquentDb.Close()
	if err != nil {
		log.Fatalln("dial mysql error", err)
	}

	// 初始化日志
	logger.New(&config.AliLog)

	// http服务器
	gin.DisableConsoleColor()
	app := gin.New()
	app.Use(gin.Recovery())

	if conf.Debug {
		gin.SetMode(gin.DebugMode)

		// 挂载中间件
		app.Use(middleware.CrossDomainForDebug())

	} else {
		gin.SetMode(gin.ReleaseMode)

		// 挂载中间件
		app.Use(middleware.CrossDomain())
	}

	// 初始化路由
	routers.Init(app)

	// 启动服务
	err = app.Run(config.Addr)
	if err != nil {
		log.Println("app run error", err)
	}
}
