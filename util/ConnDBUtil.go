package util

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnDB() *gorm.DB {
	dsn := "root:admin@tcp(127.0.0.1:3306)/gotask?charset=utf8mb4&parseTime=True&loc=Local"
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
