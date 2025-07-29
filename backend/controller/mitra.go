// controller/mitra_controller.go
package controller

import (
	"backend/interfaces"
	"backend/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type MitraController struct {
	Service interfaces.MitraService
}

func NewMitraController(service interfaces.MitraService) *MitraController {
	return &MitraController{Service: service}
}

func (c *MitraController) Create(ctx echo.Context) error {
	var data models.MitraCreate
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "Invalid input"})
	}
	if err := ctx.Validate(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "Validation failed", Data: err.Error()})
	}
	res, err := c.Service.Create(ctx.Request().Context(), &data)
	if err != nil {
		return ctx.JSON(res.StatusCode, res)
	}
	return ctx.JSON(http.StatusCreated, res)
}

func (c *MitraController) GetByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	res, err := c.Service.GetByID(ctx.Request().Context(), uint(id))
	if err != nil {
		return ctx.JSON(res.StatusCode, res)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (c *MitraController) Update(ctx echo.Context) error {
	var data models.MitraUpdate
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "Invalid input"})
	}
	if err := ctx.Validate(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "Validation failed", Data: err.Error()})
	}
	res, err := c.Service.Update(ctx.Request().Context(), &data)
	if err != nil {
		return ctx.JSON(res.StatusCode, res)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (c *MitraController) Delete(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	res, err := c.Service.Delete(ctx.Request().Context(), uint(id))
	if err != nil {
		return ctx.JSON(res.StatusCode, res)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (c *MitraController) GetAll(ctx echo.Context) error {
	var filter models.MitraFilter
	if err := ctx.Bind(&filter); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "Invalid input"})
	}
	res, err := c.Service.GetAll(ctx.Request().Context(), &filter)
	if err != nil {
		return ctx.JSON(res.StatusCode, res)
	}
	return ctx.JSON(http.StatusOK, res)
}
