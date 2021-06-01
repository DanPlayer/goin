package db

import (
	"github.com/bluele/gcache"
	"goin/conf"
)

type Gc struct {
	Cache gcache.Cache
}

var (
	GCache *Gc = nil
)

func init() {
	GCache = &Gc{
		Cache: nil,
	}
}

func GcacheDial(config *conf.GcacheConf) error {
	cache := gcache.New(config.Size).LRU().Build()
	GCache.Cache = cache

	return nil
}
