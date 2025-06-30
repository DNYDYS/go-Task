package util

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnDB() *gorm.DB {
	dsn := "root:admin@tcp(127.0.0.1:3306)/cms?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlDb, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	db, err := mysqlDb.DB()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)
	mysqlDb.Debug()
	return mysqlDb
}

func ConnRdb() *redis.Client {
	// redis cli
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return rdb

}
