package db

import (
	"database/sql"
	"fmt"
	"goin/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	Eloquent   *gorm.DB
	EloquentDb *sql.DB
)

func MysqlDial(config *conf.MysqlConf) error {
	var err error
	Eloquent, err = gorm.Open(mysql.Open(config.DSN), &gorm.Config{})

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if Eloquent.Error != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}
	EloquentDb, err = Eloquent.DB()
	if err != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}
	EloquentDb.SetMaxIdleConns(100)
	EloquentDb.SetMaxOpenConns(5000)
	EloquentDb.SetConnMaxLifetime(time.Second * 60)
	return nil
}
