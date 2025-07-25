package controller

import (
	"backend/interfaces"
	"backend/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"fmt"
)

type AssessmentController struct {
	AssessmentServices interfaces.AssessmentService
}
func NewAssessmentController(assessmentService interfaces.AssessmentService) *AssessmentController {
	return &AssessmentController{
		AssessmentServices: assessmentService,
	}
}


func (c *AssessmentController) CreateAssessment(ctx echo.Context) error {
	var assessment models.Assessment
	if err := ctx.Bind(&assessment); err != nil {
		fmt.Println("Bind error:", err)
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	if err := c.AssessmentServices.Create(&assessment); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create assessment"})
	}
	return ctx.JSON(http.StatusCreated, assessment)
}

func (c *AssessmentController) GetAssessmentByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid ID"})
	}
	assessment, err := c.AssessmentServices.GetByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Assessment not found"})
	}
	return ctx.JSON(http.StatusOK, assessment)
}

func (c *AssessmentController) UpdateAssessment(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid ID"})
	}
	var assessment models.Assessment
	if err := ctx.Bind(&assessment); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	assessment.ID = uint(id)

	if err := c.AssessmentServices.Update(&assessment); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update assessment"})
	}
	return ctx.JSON(http.StatusOK, assessment)
}

func (c *AssessmentController) DeleteAssessment(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid ID"})
	}
	if err := c.AssessmentServices.Delete(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete assessment"})
	}
	return ctx.String(http.StatusOK, "Assessment deleted successfully")
}
