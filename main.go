package main

import (
	"belajar-gin/db"
	"belajar-gin/server"
	"belajar-gin/server/controller"
	"belajar-gin/server/repository"
	"belajar-gin/server/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepo(db)
	userService := service.NewService(userRepo)
	userHandler := controller.NewUserHandler(userService)

	router := gin.Default()
	app := server.NewRouter(router, userHandler)

	app.Start(":8080")
}
