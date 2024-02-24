package model

import (
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbShareInstance *gorm.DB
	once            sync.Once
)

func DBShareInstrance() *gorm.DB {
	// 连接到 MySQL 数据库
	dsn := "user:password@tcp(localhost:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
	once.Do(func() {
		_dbShareInstance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("open mysql dataBase error!!!")
		}
		dbShareInstance = _dbShareInstance
	})
	return dbShareInstance
}

func Gcd(a int, b int) int {
	return a + b
}
