package application

import (
	"github.com/gin-gonic/gin"
	"goin/conf"
	"goin/middleware"
	helloV1 "goin/routers/application/hello/v1"
	"goin/utils"
)

func Hello(c *gin.Context) {
	utils.OutJsonOk(c, "这里是应用层服务，你好!")
}

func Init(e *gin.Engine, config *conf.AppConf) {
	hello := e.Group("/hello")
	{
		hello.GET("/", Hello)
		v1 := hello.Group("/v1")

		v1.GET("/index", middleware.MemberAuth(config.Auth.Token), helloV1.Hello)
	}
}
