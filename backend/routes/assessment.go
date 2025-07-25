package routes

import (
	"backend/controller"
	"backend/services"
	"net/http"
	"database/sql"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Berhasil terkoneksi dengan database!")
	})
	
	

	return e
}

func RegisterAssessmentRoutes(e *echo.Echo, db *sql.DB) {
	assessmentService := services.NewAssessmentService(db)
	assessmentController := controller.NewAssessmentController(assessmentService)

	e.POST("/assessment", assessmentController.CreateAssessment)
	e.GET("/assessment/:id", assessmentController.GetAssessmentByID)
	e.PUT("/assessment/:id", assessmentController.UpdateAssessment)
	e.DELETE("/assessment/:id", assessmentController.DeleteAssessment)
}