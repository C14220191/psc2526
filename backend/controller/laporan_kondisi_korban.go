package controller

import (
	"backend/interfaces"
	"backend/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type LaporanKondisiKorbanController struct {
	LaporanKondisiKorbanServices interfaces.LaporanKondisiKorbanInterface
}

func NewLaporanKondisiKorbanController(service interfaces.LaporanKondisiKorbanInterface) *LaporanKondisiKorbanController {
	return &LaporanKondisiKorbanController{
		LaporanKondisiKorbanServices: service,
	}
}

func (c *LaporanKondisiKorbanController) Create(ctx echo.Context) error {
	var data models.LaporanKondisiKorban
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid input",
		})
	}
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	result, err := c.LaporanKondisiKorbanServices.Create(ctx.Request().Context(), &data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	return ctx.JSON(result.StatusCode, result)
}

func (c *LaporanKondisiKorbanController) GetByID(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
		})
	}
	var laporan models.LaporanKondisiKorban
	result, err := c.LaporanKondisiKorbanServices.GetByID(ctx.Request().Context(), &laporan, uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, result)
	}
	return ctx.JSON(result.StatusCode, result)
}

func (c *LaporanKondisiKorbanController) Update(ctx echo.Context) error {
	var data models.LaporanKondisiKorban
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid input",
		})
	}
	data.UpdatedAt = time.Now()
	result, err := c.LaporanKondisiKorbanServices.Update(ctx.Request().Context(), &data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	return ctx.JSON(result.StatusCode, result)
}

func (c *LaporanKondisiKorbanController) Delete(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
		})
	}
	err = c.LaporanKondisiKorbanServices.Delete(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to delete laporan",
		})
	}
	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Laporan deleted successfully",
	})
}
