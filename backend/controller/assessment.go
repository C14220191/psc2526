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
	AssessmentServices interfaces.AssessmentInterface
}
func NewAssessmentController(assessmentService interfaces.AssessmentInterface) *AssessmentController {
	return &AssessmentController{
		AssessmentServices: assessmentService,
	}
}


func (c *AssessmentController) CreateAssessment(ctx echo.Context) error {
	var assessment models.AssessmentCreate
	if err := ctx.Bind(&assessment); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, models.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "unprocessable request",
			Data:       nil,
		})
	}

	err := ctx.Validate(&assessment)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Validation failed: %v", err),
			Data:       nil,
		})
	}
	
	response, err := c.AssessmentServices.Create(&assessment, ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to create assessment",
			Data:       nil,
		})
	}
	return ctx.JSON(http.StatusCreated, response)
}
func (c *AssessmentController) GetAllAssessments(ctx echo.Context) error {
	var data models.AssessmentGetAllResponse
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request",
			Data:       nil,
		})
	}
	response, err := c.AssessmentServices.GetAll(ctx.Request().Context(), data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve assessments",
			Data:       nil,
		})
	}
	return ctx.JSON(http.StatusOK, response)
}
func (c *AssessmentController) GetAssessmentByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid ID"})
	}

	assessment := models.Assessment{}
	response, err := c.AssessmentServices.GetByID(ctx.Request().Context(), assessment, uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Assessment not found"})
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *AssessmentController) UpdateAssessment(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid ID"})
	}
	var assessment models.AssessmentUpdate
	assessment.ID = uint(id)

	if err := ctx.Bind(&assessment); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}
	if err := ctx.Validate(&assessment); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, echo.Map{"error": fmt.Sprintf("Validation failed: %v", err)})
	}
	response, err := c.AssessmentServices.Update(&assessment, ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update assessment"})
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *AssessmentController) DeleteAssessment(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid ID"})
	}
	if _,err := c.AssessmentServices.Delete(uint(id), ctx.Request().Context()); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete assessment"})
	}
	return ctx.String(http.StatusOK, "Assessment deleted successfully")
}
