package routes

import (
	"backend/controller"
	"backend/services"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func DokumentasiMitraRoute(e *echo.Echo, db *sql.DB) error {
	service := services.NewDokumentasiMitraService(db)
	controller := controller.NewDokumentasiMitraController(service)

	e.POST("/dokumentasi-mitra", controller.Create)
	e.GET("/dokumentasi-mitra/:id", controller.GetByID)
	e.PUT("/dokumentasi-mitra/:id", controller.Update)
	e.DELETE("/dokumentasi-mitra/:id", controller.Delete)

	return nil
}
