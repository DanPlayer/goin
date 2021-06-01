package conf

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

var Debug = true
var confPath = "app.dev.yaml"

// Init 初始化配置文件路径
func Init() {
	RootPath := filepath.Dir(os.Args[0])
	env := flag.String("env", "dev", "-env=dev")
	flag.Parse()
	if env == nil {
		return
	}

	// 配置文件
	tmp := filepath.Join(RootPath, "app."+*env+".yaml")
	if _, err := os.Stat(tmp); err == nil {
		confPath = tmp
	}

	// 输出配置文件
	log.Println("config file:", confPath)

	// 环境
	if *env == "prod" {
		Debug = false
	}
}

func GetConfPath() string {
	return confPath
}
