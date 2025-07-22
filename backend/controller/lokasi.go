package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type LokasiController struct {
	LokasiService interfaces.LokasiService
}

func (c *LokasiController) CreateLokasi(w http.ResponseWriter, r *http.Request) {
	var lokasi models.Lokasi
	if err := json.NewDecoder(r.Body).Decode(&lokasi); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.LokasiService.Create(&lokasi); err != nil {
		http.Error(w, "Failed to create lokasi", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(lokasi)
}

func (c *LokasiController) GetLokasiByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	lokasi, err := c.LokasiService.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Lokasi not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lokasi)
}

func (c *LokasiController) UpdateLokasi(w http.ResponseWriter, r *http.Request) {
	var lokasi models.Lokasi
	if err := json.NewDecoder(r.Body).Decode(&lokasi); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.LokasiService.Update(&lokasi); err != nil {
		http.Error(w, "Failed to update lokasi", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lokasi)
}

func (c *LokasiController) DeleteLokasi(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.LokasiService.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete admin", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Lokasi deleted successfully"))
}
