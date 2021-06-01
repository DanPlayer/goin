package tests

import (
	"fmt"
	"goin/conf"
	"goin/db"
	"goin/rdb"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("初始化环境")

	// 初始化配置文件
	cw, err := os.Getwd()
	if err != nil {
		log.Fatalln("get current directory error", err)
	}
	err = conf.LoadConf(filepath.Join(cw, "../app.dev.yaml"))
	if err != nil {
		log.Fatalln("load config error", err)
	}

	// 取配置
	config := conf.GetConf()

	//初始化Mysql数据库
	err = db.MysqlDial(&config.Mysql)
	defer db.EloquentDb.Close()
	if err != nil {
		log.Fatalln("dial mysql error", err)
	}

	// 启动测试
	m.Run()
}

func TestHello(m *testing.T) {
	fmt.Printf("Hello world")
}

func TestLogin(t *testing.T) {
	rdb.Set("ke111", "121", 0)
}
