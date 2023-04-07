package main

import (
	"github.com/alireza-dehghan-nayeri/information-security-go-api/api/controller"
	"github.com/alireza-dehghan-nayeri/information-security-go-api/api/repository"
	"github.com/alireza-dehghan-nayeri/information-security-go-api/api/routes"
	"github.com/alireza-dehghan-nayeri/information-security-go-api/api/service"
	"github.com/alireza-dehghan-nayeri/information-security-go-api/infrastructure"
	"github.com/alireza-dehghan-nayeri/information-security-go-api/models"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {

	router := infrastructure.NewGinRouter()
	db := infrastructure.NewDatabase()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController, router)
	userRoute.Setup()

	db.DB.AutoMigrate(&models.User{})

	router.Gin.Run(":8000")
}
