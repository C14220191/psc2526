package routes

import (
	"backend/controller"
	"backend/services"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func KategoriKasusRoute(e *echo.Echo, db *sql.DB) error {
	service := services.NewKategoriKasusService(db)
	controller := controller.NewKategoriKasusController(service)

	e.GET("/kategori-kasus", controller.GetAll)             // opsional
	e.GET("/kategori-kasus/:id", controller.GetByID)        // opsional
	e.POST("/kategori-kasus", controller.Create)            // "add"
	e.DELETE("/kategori-kasus/:id", controller.Delete)      // "hapus"

	return nil
}
