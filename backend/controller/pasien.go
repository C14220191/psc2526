package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type PasienController struct {
	PasienServices interfaces.PasienService
}

func (c *PasienController) CreatePasien(w http.ResponseWriter, r *http.Request) {
	var pasien models.Pasien
	if err := json.NewDecoder(r.Body).Decode(&pasien); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.PasienServices.Create(&pasien); err != nil {
		http.Error(w, "Failed to create pasien", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pasien)
}

func (c *PasienController) GetPasienByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	pasien, err := c.PasienServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Pasien not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pasien)
}

func (c *PasienController) UpdatePasien(w http.ResponseWriter, r *http.Request) {
	var pasien models.Pasien
	if err := json.NewDecoder(r.Body).Decode(&pasien); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.PasienServices.Update(&pasien); err != nil {
		http.Error(w, "Failed to update pasien", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pasien)
}

func (c *PasienController) DeletePasien(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.PasienServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete pasien", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pasien deleted successfully"))
}
