package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goin/conf"
	"sync"
	"time"
)

type DB struct {
	database *mongo.Database
	lock     sync.RWMutex
}

var (
	db *DB = nil
)

func init() {
	db = &DB{
		database: nil,
		lock:     sync.RWMutex{},
	}
}

func MongoDial(config *conf.MongoConf) error {
	db.lock.RLock()
	if db.database != nil {
		db.lock.RUnlock()
		return nil
	}
	db.lock.RUnlock()

	// 链接数据库
	db.lock.Lock()
	defer db.lock.Unlock()
	client, err := mongo.NewClient(options.Client().ApplyURI(config.DSN))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	db.database = client.Database(config.DB)
	return nil
}

func Exec(collection string, f func(c *mongo.Collection)) {
	f(db.database.Collection(collection))
}
