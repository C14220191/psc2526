package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)
type VerifikasiUserController struct {
	VerifikasiUserServices interfaces.VerifikasiUserService
}

func (c *VerifikasiUserController) CreateVerifikasiUser(w http.ResponseWriter, r *http.Request) {
	var verifikasiUser models.VerifikasiUser
	if err := json.NewDecoder(r.Body).Decode(&verifikasiUser); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.VerifikasiUserServices.Create(&verifikasiUser); err != nil {
		http.Error(w, "Failed to create verifikasi user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(verifikasiUser)
}

func (c *VerifikasiUserController) GetVerifikasiUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	verifikasiUser, err := c.VerifikasiUserServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Verifikasi user not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(verifikasiUser)
}

func (c *VerifikasiUserController) UpdateVerifikasiUser(w http.ResponseWriter, r *http.Request) {
	var verifikasiUser models.VerifikasiUser
	if err := json.NewDecoder(r.Body).Decode(&verifikasiUser); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.VerifikasiUserServices.Update(&verifikasiUser); err != nil {
		http.Error(w, "Failed to update verifikasi user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(verifikasiUser)
}

func (c *VerifikasiUserController) DeleteVerifikasiUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.VerifikasiUserServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete verifikasi user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Verifikasi user deleted successfully"))
}