package routes

import (
	"backend/controller"
	"backend/services"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func DokterRoute(e *echo.Echo, db *sql.DB) error {
	dokterService := services.NewDokterService(db)
	dokterController := controller.NewDokterController(dokterService)

	// GET
	e.GET("/dokter", dokterController.GetAll)
	e.GET("/dokter/:id", dokterController.GetByID)

	// POST
	e.POST("/dokter", dokterController.Create)

	// PUT
	e.PUT("/dokter/:id", dokterController.Update)

	// DELETE
	e.DELETE("/dokter/:id", dokterController.Delete)
	return nil
}
