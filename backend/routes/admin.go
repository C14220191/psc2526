package routes

import (
	"backend/controller"
	"backend/services"
	"database/sql"

	"github.com/labstack/echo/v4"
)

// /api/v1
func AdminRoute(e *echo.Echo, db *sql.DB) error {
	adminService := services.NewAdminService(db)
	adminController := controller.NewAdminController(adminService)
	e.GET("/admin", adminController.GetAll)
	e.POST("/admin", adminController.Create)
	e.DELETE("admin/:id", adminController.DeleteAdmin)
	return nil
}