// routes/mitra_route.go
package routes

import (
	"backend/controller"
	"backend/services"
	"database/sql"
	"github.com/labstack/echo/v4"
)

func InitMitraRoute(e *echo.Echo, db *sql.DB) {
	mitraService := services.NewMitraService(db)
	mitraController := controller.NewMitraController(mitraService)

	e.POST("/mitra", mitraController.Create)
	e.GET("/mitra/:id", mitraController.GetByID)
	e.PUT("/mitra/:id", mitraController.Update)
	e.DELETE("/mitra/:id", mitraController.Delete)
	e.GET("/mitra", mitraController.GetAll)
}
