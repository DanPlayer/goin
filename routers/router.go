package routers

import (
	"github.com/gin-gonic/gin"
	"goin/conf"
	"goin/routers/application"
	"goin/utils"
)

func Hello(c *gin.Context) {
	utils.OutJsonOk(c, "这里是后端服务，你好!")
}

func Init(e *gin.Engine) {
	config := conf.GetConf()

	e.GET("/", Hello)

	//application接口
	application.Init(e, config)
}
