package controller

import (
	"backend/interfaces"
	"backend/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type DokumentasiMitraController struct {
	DokumentasiMitraServices interfaces.DokumentasiMitraInterface
}

func NewDokumentasiMitraController(service interfaces.DokumentasiMitraInterface) *DokumentasiMitraController {
	return &DokumentasiMitraController{
		DokumentasiMitraServices: service,
	}
}

func (c *DokumentasiMitraController) Create(ctx echo.Context) error {
	var data models.DokumentasiMitra
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid input",
		})
	}
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	result, err := c.DokumentasiMitraServices.Create(ctx.Request().Context(), &data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	return ctx.JSON(result.StatusCode, result)
}

func (c *DokumentasiMitraController) GetByID(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
		})
	}
	var dokumentasi models.DokumentasiMitra
	result, err := c.DokumentasiMitraServices.GetByID(ctx.Request().Context(), &dokumentasi, uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, result)
	}
	return ctx.JSON(result.StatusCode, result)
}

func (c *DokumentasiMitraController) Update(ctx echo.Context) error {
	var data models.DokumentasiMitra
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid input",
		})
	}
	data.UpdatedAt = time.Now()
	result, err := c.DokumentasiMitraServices.Update(ctx.Request().Context(), &data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	return ctx.JSON(result.StatusCode, result)
}

func (c *DokumentasiMitraController) Delete(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
		})
	}
	err = c.DokumentasiMitraServices.Delete(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to delete dokumentasi mitra",
		})
	}
	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Dokumentasi mitra deleted successfully",
	})
}
