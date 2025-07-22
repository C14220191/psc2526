package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type BeritaController struct {
	BeritaServices interfaces.BeritaService
}

func (c *BeritaController) CreateBerita(w http.ResponseWriter, r *http.Request) {
	var berita models.Berita
	if err := json.NewDecoder(r.Body).Decode(&berita); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.BeritaServices.Create(&berita); err != nil {
		http.Error(w, "Failed to create berita", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(berita)
}

func (c *BeritaController) GetBeritaByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	berita, err := c.BeritaServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Berita not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(berita)
}

func (c *BeritaController) UpdateBerita(w http.ResponseWriter, r *http.Request) {
	var berita models.Berita
	if err := json.NewDecoder(r.Body).Decode(&berita); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.BeritaServices.Update(&berita); err != nil {
		http.Error(w, "Failed to update berita", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(berita)
}

func (c *BeritaController) DeleteBerita(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.BeritaServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete berita", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Berita deleted successfully"))
}
