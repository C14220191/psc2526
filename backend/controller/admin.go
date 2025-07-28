package controller

import (
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	AdminServices interfaces.AdminInterface
}

func NewAdminController(adminServices interfaces.AdminInterface) *AdminController {
	return &AdminController{
		AdminServices: adminServices,
	}
}

func (cc *AdminController) GetAll(c echo.Context) error {
	var filter models.AdminFilter

	if err := c.Bind(&filter); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "unprocessable request",
			Data:       nil,
		})
	}

	if err := c.Validate(&filter); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"message":     "Validation failed",
		})
	}
	result, err := cc.AdminServices.GetAll(c.Request().Context(), &filter)
	if err != nil {
		return echo.NewHTTPError(result.StatusCode, result)
	}
	return c.JSON(result.StatusCode, result)
}

func (cc *AdminController) Create(c echo.Context) error {
	var data models.AdminCreate
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "unprocessable request",
			Data:       nil,
		})
	}

	if err := c.Validate(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"message":     "Validation failed",
			"errors":      err.Error(),
		})
	}

	result, err := cc.AdminServices.Create(c.Request().Context(), &data)
	if err != nil {
		return echo.NewHTTPError(result.StatusCode, result)
	}

	return c.JSON(result.StatusCode, result)
}

func (cc *AdminController) GetByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	admin := &models.Admin{}
	result, err := cc.AdminServices.GetByID(c.Request().Context(), admin, uint(id))
	if err != nil {
		return echo.NewHTTPError(result.StatusCode, result)
	}

	return c.JSON(result.StatusCode, result)
}

func (cc *AdminController) Update(c echo.Context) error {
	var adminUpdate models.AdminUpdate

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID parameter",
			Data:       nil,
		})
	}
	adminUpdate.ID = uint(id) 

	if err := c.Bind(&adminUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request body",
			Data:       nil,
		})
	}

	if err := c.Validate(&adminUpdate); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Validation failed",
			Data:       nil,
		})
	}

	res, err := cc.AdminServices.Update(c.Request().Context(), &adminUpdate)
	if err != nil {
		return c.JSON(res.StatusCode, res)
	}

	return c.JSON(http.StatusOK, res)
}

func (cc *AdminController) DeleteAdmin(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	err = cc.AdminServices.Delete(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete admin")
	}

	return c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Admin deleted successfully",
		Data:       nil,
	})
}
