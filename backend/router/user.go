package router

import (
	"backend/controller"
	"backend/services"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(e *echo.Echo, db *sql.DB) {
	userService := services.NewUserService(db)
	userController := controller.NewUserController(userService)

	e.POST("/user", userController.CreateUser)
	e.GET("/user/:id", userController.GetUserByID)
	e.PUT("/user/:id", userController.UpdateUser)
	e.DELETE("/user/:id", userController.DeleteUser)
}
