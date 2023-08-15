package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connected to database success")

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	regist := router.Group("/api/v1")

	regist.POST("/users", userHandler.RegisterUser)
	regist.POST("/login", userHandler.Login)

	router.Run()
}
