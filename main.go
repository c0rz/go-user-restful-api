package main

import (
	"log"
	"simple-api-go-c0rz/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	usersModels := user.ConnectDB(db)
	userSerivce := user.NewService(usersModels)
	userHandler := user.NewUserHandler(userSerivce)

	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")

	api.GET("/")

	router.Run(":8080")
}
