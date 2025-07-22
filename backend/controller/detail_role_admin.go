package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type DetailRoleAdminController struct {
    DetailRoleAdminServices interfaces.DetailRoleAdminService
}

func (c *DetailRoleAdminController) CreateDetailRoleAdmin(w http.ResponseWriter, r *http.Request) {
    var detailRoleAdmin models.DetailRoleAdmin
	if err := json.NewDecoder(r.Body).Decode(&detailRoleAdmin); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.DetailRoleAdminServices.Create(&detailRoleAdmin); err != nil {
		http.Error(w, "Failed to create detail role admin", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(detailRoleAdmin)
}

func (c *DetailRoleAdminController) GetDetailRoleAdminByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	detailRoleAdmin, err := c.DetailRoleAdminServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Detail role admin not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(detailRoleAdmin)
}

func (c *DetailRoleAdminController) UpdateDetailRoleAdmin(w http.ResponseWriter, r *http.Request) {
	var detailRoleAdmin models.DetailRoleAdmin
	if err := json.NewDecoder(r.Body).Decode(&detailRoleAdmin); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.DetailRoleAdminServices.Update(&detailRoleAdmin); err != nil {
		http.Error(w, "Failed to update detail role admin", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(detailRoleAdmin)
}

func (c *DetailRoleAdminController) DeleteDetailRoleAdmin(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.DetailRoleAdminServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete detail role admin", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Detail role admin deleted successfully"))
}
