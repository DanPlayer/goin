package v1

import (
	"github.com/gin-gonic/gin"
	"goin/service"
)

func Hello(c *gin.Context) {
	token := c.Query("token")
	service.HelloService.Hello()
	println(token)
}
