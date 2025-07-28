package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"

	"github.com/labstack/echo/v4"
)

type BeritaController struct {
	BeritaInterfaces interfaces.BeritaInterface
}

func NewBeritaController(beritaInterfaces interfaces.BeritaInterface) *BeritaController {
	return &BeritaController{BeritaInterfaces: beritaInterfaces}
}

func (cc *BeritaController) CreateBerita(c echo.Context) error {
	var berita models.BeritaCreate
	if err := c.Bind(&berita); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request",
			Data:       nil,
		})
	}

	err := c.Validate(&berita)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Validation failed",
			Data:       nil,
		})
	}

	response, err := cc.BeritaInterfaces.Create(&berita, c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": "Failed to create berita",
			"data":    nil,
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, response)
}

func (cc *BeritaController) GetBeritaByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
			Data:       nil,
		})
	}

	response, err := cc.BeritaInterfaces.GetByID(uint(id), c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound, response)
	}

	return c.JSON(http.StatusOK, response)
}

func (cc *BeritaController) UpdateBerita(c echo.Context) error {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "Invalid ID",
			"data":    nil,
		})
		return err
	}

	fmt.Println("berhasil ambil id")
	var berita models.BeritaUpdate
	fmt.Println("berhasil bind berita")
	if err := c.Bind(&berita); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "Invalid request",
			"data":    nil,
			"error":   err.Error(),
		})
		return err
	}

	err = c.Validate(&berita)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "Validation failed",
			"data":    nil,
			"error":   err.Error(),
		})
		return err
	}

	response, err := cc.BeritaInterfaces.Update(&berita, c.Request().Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": "Failed to update berita",
			"data":    nil,
			"error":   err.Error(),
		})
		return err
	}
	return c.JSON(http.StatusOK, response)
}

func (cc *BeritaController) DeleteBerita(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "Invalid ID",
			"data":    nil,
		})
	}

	response, err := cc.BeritaInterfaces.Delete(uint(id), c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": "Failed to delete berita",
			"data":    nil,
			"error":   err.Error(),
		})
	}

	response.StatusCode = http.StatusOK
	response.Message = "Berita deleted successfully"
	response.Data = nil

	return c.JSON(http.StatusOK, response)
}
