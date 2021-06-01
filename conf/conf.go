package conf

import (
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"
)

// MysqlConf mysql配置
type MysqlConf struct {
	DSN string `yaml:"dsn"`
}

// RedisConf redis配置
type RedisConf struct {
	Addr     string `yaml:"addr"`
	DB       int    `yaml:"db"`
	Password string `yaml:"password"`
}

// MongoConf mongodb 配置
type MongoConf struct {
	DB  string `yaml:"db"`
	DSN string `yaml:"dsn"`
}

// GcacheConf gcache配置
type GcacheConf struct {
	Size int `yaml:"size"`
}

// AuthConf auth 配置
type AuthConf struct {
	AdmSecret    string `yaml:"admin_secret"`
	MemberSecret string `yaml:"member_secret"`
	Token        string `yaml:"token"`
}

// QiniuConf 七牛配置
type QiniuConf struct {
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	Bucket    string `yaml:"bucket"`
	Domain    string `yaml:"domain"`
}

type AliLogConf struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
}

// AppConf app配置
type AppConf struct {
	Addr    string            `yaml:"addr"`
	AppName string            `yaml:"appname"`
	DocAuth map[string]string `yaml:"doc_auth"`
	Auth    AuthConf          `yaml:"auth"`
	Mysql   MysqlConf         `yaml:"mysql"`
	Redis   RedisConf         `yaml:"redis"`
	Mongo   MongoConf         `yaml:"mongo"`
	Gcache  GcacheConf        `yaml:"gcache"`
	Qiniu   QiniuConf         `yaml:"qiniu"`
	AliLog  AliLogConf        `yaml:"ali_log"`
}

var _conf *AppConf = nil
var _lock = sync.RWMutex{}

func LoadConf(ConfPath string) error {
	_lock.RLock()
	if _conf != nil {
		_lock.RUnlock()
		return nil
	}
	_lock.RUnlock()

	_lock.Lock()
	defer _lock.Unlock()
	_conf = &AppConf{}
	data, err := ioutil.ReadFile(ConfPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &_conf)
	return err
}

func GetConf() *AppConf {
	return _conf
}
