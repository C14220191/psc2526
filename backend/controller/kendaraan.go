package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type KendaraanController struct {
	KendaraanServices interfaces.KendaraanService
}

func (c *KendaraanController) CreateKendaraan(w http.ResponseWriter, r *http.Request) {
	var kendaraan models.Kendaraan
	if err := json.NewDecoder(r.Body).Decode(&kendaraan); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.KendaraanServices.Create(&kendaraan); err != nil {
		http.Error(w, "Failed to create kendaraan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(kendaraan)
}

func (c *KendaraanController) GetKendaraanByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	kendaraan, err := c.KendaraanServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Kendaraan not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(kendaraan)
}

func (c *KendaraanController) UpdateKendaraan(w http.ResponseWriter, r *http.Request) {
	var kendaraan models.Kendaraan
	if err := json.NewDecoder(r.Body).Decode(&kendaraan); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.KendaraanServices.Update(&kendaraan); err != nil {
		http.Error(w, "Failed to update kendaraan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(kendaraan)
}

func (c *KendaraanController) DeleteKendaraan(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.KendaraanServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete kendaraan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Kendaraan deleted successfully"))
}
