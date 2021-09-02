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
	connect := "root:@tcp(127.0.0.1:3306)/golangdatabase?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	usersModels := user.ConnectDB(db)
	userSerivces := user.NewService(usersModels)
	userControllers := user.Controllers(userSerivces)

	router := gin.Default()
	router.NoRoute(userControllers.NotFound)
	router.Use(cors.Default())

	api := router.Group("/api/v1")

	api.GET("/list", userControllers.GetUsers)
	api.POST("/create", userControllers.RegisterUsers)
	api.PATCH("/update/:id", userControllers.UpdateUsers)
	api.DELETE("/delete/:id", userControllers.DeleteUser)

	router.Run(":8080")
}
