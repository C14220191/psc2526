package routes

import (
	"backend/controller"
	"backend/services"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func LaporanKondisiKorbanRoute(e *echo.Echo, db *sql.DB) error {
	service := services.NewLaporanKondisiKorbanService(db)
	controller := controller.NewLaporanKondisiKorbanController(service)

	e.POST("/laporan-kondisi-korban", controller.Create)
	e.GET("/laporan-kondisi-korban/:id", controller.GetByID)
	e.PUT("/laporan-kondisi-korban/:id", controller.Update)
	e.DELETE("/laporan-kondisi-korban/:id", controller.Delete)

	return nil
}
