package routes

import (
	"github.com/alireza-dehghan-nayeri/information-security-go-api/api/controller"
	"github.com/alireza-dehghan-nayeri/information-security-go-api/api/middlewares"
	"github.com/alireza-dehghan-nayeri/information-security-go-api/infrastructure"
)

// UserRoute -> Route for user module
type UserRoute struct {
	Handler    infrastructure.GinRouter
	Controller controller.UserController
}

// NewUserRoute -> initializes new instance of UserRoute
func NewUserRoute(
	controller controller.UserController,
	handler infrastructure.GinRouter,
) UserRoute {
	return UserRoute{
		Handler:    handler,
		Controller: controller,
	}
}

// Setup -> setups user routes
func (u UserRoute) Setup() {
	user := u.Handler.Gin.Group("/auth")
	{
		user.Use(middlewares.RateLimiterMiddleware())
		user.POST("/register", u.Controller.CreateUser)
		user.POST("/login", u.Controller.LoginUser)
	}
}
