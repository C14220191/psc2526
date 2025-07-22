package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type DetailPrivilegeRoleController struct {
	DetailPrivilegeRoleServices interfaces.DetailPrivilegeRoleService
}

func (c *DetailPrivilegeRoleController) CreateDetailPrivilegeRole(w http.ResponseWriter, r *http.Request) {
	var detailPrivilegeRole models.DetailPrivilegeRole
	if err := json.NewDecoder(r.Body).Decode(&detailPrivilegeRole); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.DetailPrivilegeRoleServices.Create(&detailPrivilegeRole); err != nil {
		http.Error(w, "Failed to create detail privilege role", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(detailPrivilegeRole)
}

func (c *DetailPrivilegeRoleController) GetDetailPrivilegeRoleByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	detailPrivilegeRole, err := c.DetailPrivilegeRoleServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Detail privilege role not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(detailPrivilegeRole)
}

func (c *DetailPrivilegeRoleController) UpdateDetailPrivilegeRole(w http.ResponseWriter, r *http.Request) {
	var detailPrivilegeRole models.DetailPrivilegeRole
	if err := json.NewDecoder(r.Body).Decode(&detailPrivilegeRole); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.DetailPrivilegeRoleServices.Update(&detailPrivilegeRole); err != nil {
		http.Error(w, "Failed to update detail privilege role", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(detailPrivilegeRole)
}

func (c *DetailPrivilegeRoleController) DeleteDetailPrivilegeRole(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.DetailPrivilegeRoleServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete detail privilege role", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Detail privilege role deleted successfully"))
}
