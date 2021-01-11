package database

import (
	"fmt"
	"test/server/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB for connecting on other package
var DB *gorm.DB

// InitDatabase for init database in the start
func InitDatabase() {
	var err error
	dsn := "root:@tcp(127.0.0.1)/test_server?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("database is error")
	}

	DB.AutoMigrate(&model.User{})

	fmt.Println("Database Connected")
}
