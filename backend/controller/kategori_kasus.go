package controller

import (
	"backend/interfaces"
	"backend/models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type KategoriKasusController struct {
	KategoriKasusService interfaces.KategoriKasusInterface
}

func NewKategoriKasusController(s interfaces.KategoriKasusInterface) *KategoriKasusController {
	return &KategoriKasusController{KategoriKasusService: s}
}

func (c *KategoriKasusController) Create(ctx echo.Context) error {
	var data models.KategoriKasus
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid input")
	}
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	if err := c.KategoriKasusService.Create(&data); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to create kategori kasus")
	}
	return ctx.JSON(http.StatusCreated, data)
}

func (c *KategoriKasusController) Delete(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid ID")
	}
	if err := c.KategoriKasusService.Delete(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to delete kategori kasus")
	}
	return ctx.JSON(http.StatusOK, "Kategori kasus deleted successfully")
}

// Optional: Get All
func (c *KategoriKasusController) GetAll(ctx echo.Context) error {
	// dummy filterless fetch (you can enhance with query param filter later)
	list, err := c.KategoriKasusService.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to fetch data")
	}
	return ctx.JSON(http.StatusOK, list)
}

// Optional: GetByID
func (c *KategoriKasusController) GetByID(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid ID")
	}
	result, err := c.KategoriKasusService.GetByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Data not found")
	}
	return ctx.JSON(http.StatusOK, result)
}
