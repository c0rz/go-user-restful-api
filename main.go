package main

import (
	"fmt"
	"log"
	"simple-api-go-c0rz/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	connect := "root:@tcp(127.0.0.1:3306)/belajar?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	_ = user.ConnectDB(db)

	fmt.Println("Connected")
}
