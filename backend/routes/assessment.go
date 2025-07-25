package routes
import (
	"net/http"
	"github.com/labstack/echo/v4"
	"backend/controller"
)

func Init() *echo.Echo {
	e := echo.New()
	
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Berhasil terkoneksi dengan database!")
	})
	
	

	return e
}

func RegisterAssessmentRoutes(e *echo.Echo, c *controller.AssessmentController) {
	e.POST("/assessment", c.CreateAssessment)
	e.GET("/assessment/:id", c.GetAssessmentByID)
	e.PUT("/assessment/:id", c.UpdateAssessment)
	e.DELETE("/assessment/:id", c.DeleteAssessment)
}