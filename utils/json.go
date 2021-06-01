package utils

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

const (
	DATE_FORMAT             = "2006-01-02"
	DATETIME_FORMAT         = "2006-01-02 15:04:05"
	DATETIMEWITHZONE_FORMAT = "2006-01-02 15:04:05 -07"
)

func OutJsonOk(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"rtn": 0,
		"msg": msg,
	})
}

func OutJsonErro(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"rtn": 1,
		"msg": msg,
	})
}

func ActionOutJsonErro(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"rtn": 2,
		"msg": msg,
	})
}

func OutJson(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"rtn":  0,
		"msg":  "成功",
		"data": data,
	})
}

func OutParamErrorJson(c *gin.Context) {
	c.JSON(200, gin.H{
		"rtn": 1,
		"msg": "参数错误",
	})
}

func OutErrorJson(c *gin.Context, err error) {
	c.JSON(200, gin.H{
		"rtn": 1,
		"msg": err.Error(),
	})
}

// OutAuthNeedError 需要登录
func OutAuthNeedError(c *gin.Context) {
	c.JSON(200, gin.H{
		"rtn": -1,
		"msg": "该功能需要登录!",
	})
}

// OutAuthOutdatedError 输出过期错误
func OutAuthOutdatedError(c *gin.Context) {
	c.JSON(200, gin.H{
		"rtn": -2,
		"msg": "登录态已过期,请重新登录!",
	})
}

func OutErrorJsonWithStr(c *gin.Context, err string) {
	c.JSON(200, gin.H{
		"rtn": 1,
		"msg": err,
	})
}

func OutSolveSuccessJsonWithData(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"rtn":  0,
		"msg":  "解题成功",
		"data": data,
	})
}

func OutSolveFailJsonWithData(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"rtn":  1,
		"msg":  "答案错误",
		"data": data,
	})
}

// ValidateJson 监测json数据正确性
func ValidateJson(jsonString string) (right bool, data interface{}) {
	decoder := json.NewDecoder(bytes.NewReader([]byte(jsonString)))
	decoder.UseNumber()
	err := decoder.Decode(&data)
	if err != nil {
		return false, data
	}
	return true, data
}
