package routes

import (
	"backend/controller"
	"backend/services"
	"database/sql"

	"github.com/labstack/echo/v4"
)

// /api/v1
func BeritaRoute(e *echo.Echo, db *sql.DB) error {
	beritaService := services.NewBeritaService(db)
	beritaController := controller.NewBeritaController(beritaService)

	//get
	// e.GET("/berita", beritaController.GetAll)
	e.GET("/berita/:id", beritaController.GetBeritaByID)

	//post
	e.POST("/berita", beritaController.CreateBerita)

	//put
	e.PUT("/berita/:id", beritaController.UpdateBerita)

	//delete
	e.DELETE("/berita/:id", beritaController.DeleteBerita)
	return nil
}