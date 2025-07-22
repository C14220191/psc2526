package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type PelaporController struct {
	PelaporServices interfaces.PelaporService
}

func (c *PelaporController) CreatePelapor(w http.ResponseWriter, r *http.Request) {
	var pelapor models.Pelapor
	if err := json.NewDecoder(r.Body).Decode(&pelapor); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.PelaporServices.Create(&pelapor); err != nil {
		http.Error(w, "Failed to create pelapor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pelapor)
}

func (c *PelaporController) GetPelaporByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	pelapor, err := c.PelaporServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Pelapor not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pelapor)
}

func (c *PelaporController) UpdatePelapor(w http.ResponseWriter, r *http.Request) {
	var pelapor models.Pelapor
	if err := json.NewDecoder(r.Body).Decode(&pelapor); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.PelaporServices.Update(&pelapor); err != nil {
		http.Error(w, "Failed to update pelapor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pelapor)
}

func (c *PelaporController) DeletePelapor(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.PelaporServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete pelapor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pelapor deleted successfully"))
}
