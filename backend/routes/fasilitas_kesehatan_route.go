package routes

import (
	"backend/controller"
	"backend/services"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func FasilitasKesehatanRoute(e *echo.Echo, db *sql.DB) error {
	service := services.NewFasilitasKesehatanService(db)
	controller := controller.NewFasilitasKesehatanController(service)

	e.GET("/fasilitas-kesehatan", controller.GetAll)
	e.GET("/fasilitas-kesehatan/:id", controller.GetByID)
	e.POST("/fasilitas-kesehatan", controller.Create)
	e.PUT("/fasilitas-kesehatan/:id", controller.Update)
	e.DELETE("/fasilitas-kesehatan/:id", controller.Delete)

	return nil
}
