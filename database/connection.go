package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	host     string = "localhost"
	port     string = "3306"
	username string = "root"
	password string = "12345678"
	database string = "chamsockhachhang"
)

func NewDB() *gorm.DB {
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
	)

	db, err := gorm.Open(mysql.Open(connectString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // -> Log các câu lệnh truy vấn database trong terminal
	})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}
