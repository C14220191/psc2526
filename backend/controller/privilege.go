package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type PrivilegeController struct {
	PrivilegeServices interfaces.PrivilegeService
}

func (c *PrivilegeController) CreatePrivilege(w http.ResponseWriter, r *http.Request) {
	var privilege models.Privilege
	if err := json.NewDecoder(r.Body).Decode(&privilege); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.PrivilegeServices.Create(&privilege); err != nil {
		http.Error(w, "Failed to create privilege", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(privilege)
}

func (c *PrivilegeController) GetPrivilegeByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	privilege, err := c.PrivilegeServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Privilege not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(privilege)
}

func (c *PrivilegeController) UpdatePrivilege(w http.ResponseWriter, r *http.Request) {
	var privilege models.Privilege
	if err := json.NewDecoder(r.Body).Decode(&privilege); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.PrivilegeServices.Update(&privilege); err != nil {
		http.Error(w, "Failed to update privilege", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(privilege)
}

func (c *PrivilegeController) DeletePrivilege(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.PrivilegeServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete privilege", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Privilege deleted successfully"))
}
