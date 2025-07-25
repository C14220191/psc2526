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

func (cc *AdminController) Create(c echo.Context) error {
	var data models.AdminCreate
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "unprocessable request",
			Data:       nil,
		})
	}

	if err := c.Validate(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"message":     "Validation failed",
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

func (c *AdminController) DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.AdminServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete admin", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Admin deleted successfully"))
}
