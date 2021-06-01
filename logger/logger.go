package logger

import (
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"goin/conf"
	"os"
	"time"
)

var hostname = "unknow"
var logger *producer.Producer = nil

func New(config *conf.AliLogConf)  {
	logger := makeProducer(config)
	logger.Start()

	name,err := os.Hostname()
	if err == nil {
		hostname = name
	}
}

func makeProducer(config *conf.AliLogConf) *producer.Producer{
	cfg := producer.GetDefaultProducerConfig()
	cfg.Endpoint = config.Endpoint
	cfg.AccessKeyID = config.AccessKeyID
	cfg.AccessKeySecret = config.AccessKeySecret
	return producer.InitProducer(cfg)
}

func Info(item map[string]string) error{
	logs := producer.GenerateLog(uint32(time.Now().Unix()), item)
	return logger.SendLog("puahome-web", "app_server", "icitysecret_mg", hostname, logs)
}

func LogRequest(item map[string]string) error{
	logs := producer.GenerateLog(uint32(time.Now().Unix()), item)
	return logger.SendLog("puahome-web", "app_server", "icitysecret_mg", hostname, logs)
}

