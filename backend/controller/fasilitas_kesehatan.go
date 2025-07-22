package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type FasilitasKesehatanController struct {
	FasilitasKesehatanServices interfaces.FasilitasKesehatanService
}

func (c *FasilitasKesehatanController) CreateFasilitasKesehatan(w http.ResponseWriter, r *http.Request) {
	var fasilitas models.FasilitasKesehatan
	if err := json.NewDecoder(r.Body).Decode(&fasilitas); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.FasilitasKesehatanServices.Create(&fasilitas); err != nil {
		http.Error(w, "Failed to create fasilitas kesehatan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(fasilitas)
}

func (c *FasilitasKesehatanController) GetFasilitasKesehatanByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	fasilitas, err := c.FasilitasKesehatanServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Fasilitas Kesehatan not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fasilitas)
}

func (c *FasilitasKesehatanController) UpdateFasilitasKesehatan(w http.ResponseWriter, r *http.Request) {
	var fasilitas models.FasilitasKesehatan
	if err := json.NewDecoder(r.Body).Decode(&fasilitas); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.FasilitasKesehatanServices.Update(&fasilitas); err != nil {
		http.Error(w, "Failed to update fasilitas kesehatan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fasilitas)
}

func (c *FasilitasKesehatanController) DeleteFasilitasKesehatan(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.FasilitasKesehatanServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete fasilitas kesehatan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Fasilitas Kesehatan deleted successfully"))
}
