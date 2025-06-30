package service

import (
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"task4/internal/common/logger"
)

type CmsApp struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewCmsApp() *CmsApp {
	app := &CmsApp{}
	ConnDB(app)
	ConnRdb(app)
	return app
}

func ConnDB(app *CmsApp) {
	dsn := "root:admin@tcp(127.0.0.1:3306)/cms?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlDb, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		logger.S().Debug("链接数据库异常", err)
		panic(err)
	}
	db, err := mysqlDb.DB()
	if err != nil {
		logger.S().Debug("链接数据库异常", err)
		panic(err)
	}
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)
	app.db = mysqlDb
	app.db.Debug()
}

func ConnRdb(app *CmsApp) {
	// redis cli
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logger.S().Debug("链接redis异常", err)
		panic(err)
	}
	app.rdb = rdb
}
