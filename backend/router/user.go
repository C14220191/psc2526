package router

import (
	"net/http"
	"backend/controller"
	"backend/services"
	"database/sql"
)

func RegisterUserRoutes(db *sql.DB) {
	userService := services.NewUserService(db)
	userController := &controller.UserController{UserServices: userService}

	http.HandleFunc("/user/create", userController.CreateUser)
	http.HandleFunc("/user/get", userController.GetUserByID)
	http.HandleFunc("/user/update", userController.UpdateUser)
	http.HandleFunc("/user/delete", userController.DeleteUser)
}
