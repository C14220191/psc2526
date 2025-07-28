package controller

import (
	"backend/interfaces"
	"backend/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FasilitasKesehatanController struct {
	FasilitasService interfaces.FasilitasKesehatanInterface
}

func NewFasilitasKesehatanController(s interfaces.FasilitasKesehatanInterface) *FasilitasKesehatanController {
	return &FasilitasKesehatanController{FasilitasService: s}
}

func (c *FasilitasKesehatanController) Create(ctx echo.Context) error {
	var data models.FasilitasKesehatanCreate
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid input",
			Data:       nil,
		})
	}
	result, err := c.FasilitasService.Create(ctx.Request().Context(), &data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	return ctx.JSON(result.StatusCode, result)
}

func (c *FasilitasKesehatanController) GetAll(ctx echo.Context) error {
	var filter models.FasilitasKesehatanFilter
	if err := ctx.Bind(&filter); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid filter parameters",
			Data:       nil,
		})
	}
	result, err := c.FasilitasService.GetAll(ctx.Request().Context(), &filter)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	return ctx.JSON(result.StatusCode, result)
}

func (c *FasilitasKesehatanController) GetByID(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
			Data:       nil,
		})
	}
	var fasilitas models.FasilitasKesehatan
	result, err := c.FasilitasService.GetByID(ctx.Request().Context(), &fasilitas, uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	return ctx.JSON(result.StatusCode, result)
}

func (c *FasilitasKesehatanController) Update(ctx echo.Context) error {
	var data models.FasilitasKesehatanUpdate
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid input",
			Data:       nil,
		})
	}
	result, err := c.FasilitasService.Update(ctx.Request().Context(), &data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	return ctx.JSON(result.StatusCode, result)
}

func (c *FasilitasKesehatanController) Delete(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
			Data:       nil,
		})
	}
	err = c.FasilitasService.Delete(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to delete fasilitas kesehatan",
			Data:       nil,
		})
	}
	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Fasilitas Kesehatan deleted successfully",
		Data:       nil,
	})
}
