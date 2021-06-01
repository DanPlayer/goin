package rdb

import (
	"github.com/go-redis/redis"
	"goin/conf"
	"sync"
	"time"
)

var lock sync.RWMutex
var rd *redis.Client = nil

func RedisDial(config *conf.RedisConf) error {
	lock.RLock()
	if rd != nil {
		lock.RUnlock()
		return nil
	}
	lock.RUnlock()

	lock.Lock()
	defer lock.Unlock()
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	rd = client
	return nil
}

func Set(key string, val string, expTime int32) {
	rd.Set(key, val, time.Duration(expTime)*time.Second)
}

func Get(key string) string {
	val, err := rd.Get(key).Result()
	if err != nil {
		return ""
	}
	return val
}

func Del(key string) {
	rd.Del(key)
}

func SetHash(key string, field string, value interface{}) {
	rd.HSet(key, field, value)
}

func GetHash(key string, field string) (s string, e error) {
	s, e = rd.HGet(key, field).Result()
	return
}

func DelHash(key string, field string) {
	rd.HDel(key, field)
}

func Do(Key string, Time int) {
	rd.Do("EXPIRE", Key, Time)
}
