package controller

import (
	"backend/interfaces"
	"backend/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DokterController struct {
	DokterService interfaces.DokterInterface
}

func NewDokterController(s interfaces.DokterInterface) *DokterController {
	return &DokterController{
		DokterService: s,
	}
}

func (c *DokterController) Create(ctx echo.Context) error {
	var data models.DokterCreate
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid input",
			Data:       nil,
		})
	}

	result, err := c.DokterService.Create(ctx.Request().Context(), &data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	return ctx.JSON(result.StatusCode, result)
}

func (c *DokterController) GetAll(ctx echo.Context) error {
	var filter models.DokterFilter
	if err := ctx.Bind(&filter); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid filter parameters",
			Data:       nil,
		})
	}
	result, err := c.DokterService.GetAll(ctx.Request().Context(), &filter)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	return ctx.JSON(result.StatusCode, result)
}

func (c *DokterController) GetByID(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
			Data:       nil,
		})
	}
	var dokter models.Dokter
	result, err := c.DokterService.GetByID(ctx.Request().Context(), &dokter, uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	return ctx.JSON(result.StatusCode, result)
}

func (c *DokterController) Update(ctx echo.Context) error {
	var data models.DokterUpdate
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid input",
			Data:       nil,
		})
	}
	result, err := c.DokterService.Update(ctx.Request().Context(), &data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	return ctx.JSON(result.StatusCode, result)
}

func (c *DokterController) Delete(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
			Data:       nil,
		})
	}
	err = c.DokterService.Delete(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to delete dokter",
			Data:       nil,
		})
	}
	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Dokter deleted successfully",
		Data:       nil,
	})
}
