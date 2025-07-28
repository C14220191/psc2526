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

	//get
	e.GET("/admin", adminController.GetAll)
	e.GET("/admin/:id", adminController.GetByID)

	//post
	e.POST("/admin", adminController.Create)
	
	//put
	e.PUT("/admin/:id", adminController.Update)

	//delete
	e.DELETE("/admin/:id", adminController.DeleteAdmin)
	return nil
}