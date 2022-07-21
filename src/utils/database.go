package utils

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB_HOST     = GetEnv("DB_HOST", "localhost")
	DB_NAME     = GetEnv("DB_NAME", "go-store")
	DB_PORT     = GetEnv("DB_PORT", "3306")
	DB_USERNAME = GetEnv("DB_USERNAME", "root")
	DB_PASSWORD = GetEnv("DB_PASSWORD", "root")
)

func GetConnection() *gorm.DB {
	sqlConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	db, err := gorm.Open("mysql", sqlConnection)
	if err != nil {
		panic(err)
	}
	log.Println("Database is connected")
	return db
}
