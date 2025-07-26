package controller

import (
	"encoding/json"
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
			"errors":     err.Error(),
		})
	}

	result, err := cc.AdminServices.Create(c.Request().Context(), &data)
	if err != nil {
		return echo.NewHTTPError(result.StatusCode, result)
	}

	return c.JSON(result.StatusCode, result)
}

func (c *AdminController) GetAdminByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	admin, err := c.AdminServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Admin not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(admin)
}

func (c *AdminController) UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	var admin models.Admin
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.AdminServices.Update(&admin); err != nil {
		http.Error(w, "Failed to update admin", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(admin)
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
